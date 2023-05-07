-- GENERATED CODE - DO NOT MODIFY BY HAND
drop table if exists test.my_struct;
create table if not exists test.my_struct (
    uuid uuid primary key,
    name text,
    number int,
    my_boolean boolean
    );