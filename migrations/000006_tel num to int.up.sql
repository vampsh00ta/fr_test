begin;

alter table client drop column if exists tel_number;
alter table client add  column if not exists tel_number integer;

commit;