create table if not exists users (
    id integer unsigned auto_increment primary key,
    email varchar(100) unique not null,
    password_hash varchar(100) not null,
    uuid varchar(100),
    activate_flag boolean not null,
    created_at datetime not null
);

insert into users (email, password_hash, uuid, activate_flag, created_at)
values ("gumbawcih@cidloza.eg", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "f3f69bb9-e9db-5096-a96e-f18579fe0ca2", false, now() - interval 20 minute);

insert into users (email, password_hash, uuid, activate_flag, created_at)
values ("nirjezic@ga.bm", "3e25aa8662ba1999dc949e54838e195125c229eaa6b7bacd8c634ff75f803b1d", "be896154-f725-5b13-8bd5-9934564c30d8", false, now() - interval 15 minute);

insert into users (email, password_hash, uuid, activate_flag, created_at)
values ("heera@demivde.org", "2dcced2155f0deca7bc338fed3fc57bd4d89843da25021c4a939b9117315d02f", "8f2f5682-387d-5a55-b06a-41a026a4e8c9", false, now());