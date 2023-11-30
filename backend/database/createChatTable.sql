create table if not exists chats (
    id integer primary key,
    problem_id integer unique not null,
    user_id integer unique not null,
    question text not null,
    answer text not null,
    created_at not null,
    foreign key problem_id references problems(id),
    foreign key user_id references users(id)
)