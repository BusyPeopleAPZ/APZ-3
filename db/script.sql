create database balancersdb;

create table vms(
  id serial not null,
  balancerid integer not null,
  status boolean not null default false
);

alter table machines add constraint primaryvms primary key(id);

create table balancers(
    id serial not null,
    name varchar(255) not null
);
alter table balancers add constraint primarybalancer primary key(id);
alter table machines add constraint foreignvmsbalancer foreign key(balancerid) references balancers(id);

insert into balancers(name)
values ('Alpha'),
       ('Betta'),
       ('Gamma'),
       ('Omega');

insert into vms(balancerid, status)
values (1, false),
       (2, true),
       (1, true),
       (1, true),
       (2, false),
       (3, true),
       (2, false),
       (1, true),
       (3, true),
       (3, true),
       (3, true),
       (3, true),
       (2, true),
       (3, true),
       (2, false),
       (2, true),
       (2, false),
       (4, false),
       (1, true),
       (1, false),
       (2, false),
       (2, true),
       (2, false),
       (2, false),
       (3, true),
       (4, false),
       (3, true),
       (4, false);