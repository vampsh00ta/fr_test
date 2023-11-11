begin;

create table if not exists status (
id serial primary key,
      time timestamp ,
         text varchar(255),
    message_id integer references message (id)
                    );

commit;