


create table app_reg_tab (
  id  serial primary key,
  appid  char(80) not null unique,
  appkey char(80),
  registime char(80)
);