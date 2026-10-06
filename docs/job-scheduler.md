# Job scheduler

## Flow

1. The user creates a **task**: a JS function (`executorJS`) to run.
2. The user creates a **job** for that task with a schedule: `ONCE` or `RECURRING` (cron expression).
3. When the job is saved, the execution service inserts the next **execution** into the DB as `PENDING`. If it is due within `config.execDelay.duration` (5 min), it is also pushed to the execution queue and marked `INITIATED`.
4. A cron runs every 5 minutes, fetches `PENDING` executions due in the next 5 minutes, marks them `INITIATED` and enqueues them.
5. Worker nodes pull from the queue, run the script in a goja runtime, and write the status back to the DB.

```mermaid
flowchart TD
    U[User] -->|"1. create task"| TS[Task service]
    U -->|"2. create job"| JS[Job service]
    TS --> DB[(DB<br/>tasks · jobs · executions)]
    JS -->|on save| ES[Execution service]
    ES -->|insert PENDING| DB
    ES -->|"if within 5m: enqueue"| Q[["Execution queue<br/>priority by executeAt"]]
    CR["Cron (*/5 * * * *)"] -->|"fetch PENDING, mark INITIATED"| DB
    CR -->|enqueue| Q
    W["Worker nodes 1..N<br/>goja runtime"] -->|pull| Q
    W -->|"IN_PROGRESS → SUCCESS / FAILED"| DB
```

## Execution status

No retries: a failed execution stays `FAILED`.

```mermaid
stateDiagram-v2
    [*] --> PENDING: job saved
    PENDING --> INITIATED: pushed to queue
    INITIATED --> IN_PROGRESS: worker picks it
    IN_PROGRESS --> SUCCESS: script returned
    IN_PROGRESS --> FAILED: threw or timed out
    SUCCESS --> [*]
    FAILED --> [*]
```

## Data model

```mermaid
erDiagram
    TASK ||--o{ JOB : "runs as"
    JOB ||--o{ EXECUTION : schedules

    TASK {
        uuid id PK
        string name
        text executorJS
        string createdBy
        timestamp createdAt
        timestamp updatedAt
    }
    JOB {
        uuid id PK
        uuid taskId FK
        jsonb parameters
        string type "ONCE | RECURRING"
        string expr "cron expr if RECURRING"
        string createdBy
        timestamp createdAt
        timestamp updatedAt
    }
    EXECUTION {
        uuid id PK
        uuid jobId FK
        string status "PENDING | INITIATED | IN_PROGRESS | SUCCESS | FAILED"
        timestamp executeAt
        timestamp createdAt
        timestamp updatedAt
    }
```

## Config

```yaml
execDelay:
  duration: PT5M
```

## Open questions

- **Next run of a `RECURRING` job:** saving the job creates only the first execution. Whatever marks an execution `INITIATED` (cron or execution service) should also insert the next `PENDING` one from the cron expression. Waiting for `SUCCESS`/`FAILED` would stop the job forever if a worker crashes mid-run.
