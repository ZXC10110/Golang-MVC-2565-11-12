CREATE TABLE Test.Product (
    sequence int unique,
    pid varchar(255) PRIMARY KEY,
    pname varchar(255),
    price float,
    qty int
)

INSERT INTO Test.Product VALUE (1, "Ticket-01", "Blackpink concert ticket", 79000.00, 1)
INSERT INTO Test.Product VALUE (2, "Ticket-02", "MVC ticket", 3000.00, 200)
INSERT INTO Test.Product VALUE (3, "Ticket-03", "Avenger: Endgame ticket", 1500.00, 68)
INSERT INTO Test.Product VALUE (4, "Project-A", "Bridge bar voucher", 10000.00, 100)
INSERT INTO Test.Product VALUE (5, "Project-F", "Brain junior", 50.00, 650)
INSERT INTO Test.Product VALUE (6, "10062", "Member A Class", 899.00, 2)
INSERT INTO Test.Product VALUE (7, "10063", "Member B Class", 699.00, 15)