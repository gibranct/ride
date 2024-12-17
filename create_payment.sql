drop schema if exists gct cascade;

create schema gct;

create table gct.transaction (
	transaction_id uuid primary key,
	ride_id uuid,
	amount numeric,
	date timestamp,
	status text
);