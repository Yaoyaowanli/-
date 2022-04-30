
drop table if exists t_class;
drop table if exists t_student;

create table t_class(
	classno int primary key,
	classname varchar(255)
);
create table t_student(
	no int primary key auto_increment,
	name varchar(32),
	cno int,
	foreign key(cno) references t_class(classno)
);
insert into t_class(classno,classname)values(100,'高三一班');
insert into t_class(classno,classname)values(101,'高三二班');

insert into t_student(name,cno)values('tom',100);
insert into t_student(name,cno)values('jack',100);
insert into t_student(name,cno)values('sb',101);
insert into t_student(name,cno)values('dog',100);
insert into t_student(name,cno)values('cat',101);
insert into t_student(name,cno)values('bb',100);


select * from t_class;
select * from t_student;
