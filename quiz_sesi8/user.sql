CREATE TABLE users ( user_id INT PRIMARY KEY, username VARCHAR NOT NULL UNIQUE, email VARCHAR NOT NULL UNIQUE, nama_lengkap VARCHAR NOT NULL );
create table orders ( order_id int primary key,user_id int references users(user_id),tanggal_pemesanan timestamp default now(),status varchar not null);
create table order_items (item_id int primary key,order_id int references orders(order_id),product_name varchar not null,quantity int not null,harga_per_item numeric not null);

INSERT INTO users (user_id, username, email, nama_lengkap) VALUES ( 1 , 'user_1', 'user1@example.com', 'User Satu' ),( 2 , 'user_2', 'user3@example.com', 'User Dua' ),( 3 , 'user_3', 'user3@example.com', 'User Tiga' );
insert into orders ( order_id, user_id, status ) values (1,1,'Dalam Pengiriman'),(2,2,'Selesai'),(3,3,'Dibatalkan');
 insert into order_items ( item_id, order_id, product_name, quantity, harga_per_item) values ( 1,1,'Produk A', 2, 50000), (2,1,'Produk B', 3,30000),(3,2,'Produk C',1,75000),(4,2,'Produk D',2,60000);

 select * from users;
 select * from orders;
 select * from order_items;

 select o.order_id, u.username from orders as o left join users as u on u.user_id = o.order_id;
select o.order_id, u.username, sum(oi.harga_per_item * oi.quantity) as total_harga from orders as o left join users as u on u.user_id = o.order_id right join order_items as oi on oi.order_id = o.order_id GROUP BY o.order_id, u.username;