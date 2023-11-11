begin;
create table if not exists client(
                       id serial primary key,
                       tel_number varchar(12),
                       mobile_code varchar(10),
                       tag varchar(255),
                       timezone varchar(50)
);
create table if not exists newsletter(
                           id serial primary key,
                           start_time timestamp,
                           end_time timestamp,
                           text text
);
create table if not exists  message(
                        id serial primary key,
                        send_time timestamp,
                        status  varchar(255),
                        newsletter_id integer references newsletter(id),
                        client_id integer references client(id)
);
commit;
