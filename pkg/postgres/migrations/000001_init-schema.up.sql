BEGIN;
CREATE TABLE IF NOT EXISTS "tasks"
(
    id          uuid,
    name        text        not null,
    executionJS text        not null,
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    primary key (id)
);

create table if not exists "jobs"
(
    id             uuid,
    task_id        uuid        not null,
    execution_type text        not null,
    cron_expr      text,
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now(),
    primary key (id),
    constraint fk_tasks foreign key (task_id) references tasks (id) on delete cascade
);

create table if not exists "executions"
(
    id             uuid,
    job_id         uuid        not null,
    execution_time timestamptz not null,
    status         text        not null,
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now(),
    primary key (id),
    constraint fk_jobs foreign key (job_id) references jobs (id) on delete cascade
);

create or replace function update_updated_at_column()
    returns trigger as
$$
begin
    new.updated_at = now();
    return new;
end;
$$ language plpgsql;

create trigger trigger_update_tasks_ts
    before update
    on tasks
    for each row
execute function update_updated_at_column();

create trigger trigger_update_tasks_ts
    before update
    on jobs
    for each row
execute function update_updated_at_column();

create trigger trigger_update_tasks_ts
    before update
    on executions
    for each row
execute function update_updated_at_column();
COMMIT;
