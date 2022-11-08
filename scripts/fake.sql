INSERT INTO products (id, name)
VALUES
    (1, 'C# in a Nutshell'),
    (2, 'KIA Rio'),
    (3, 'ASUS Rog Strix');

INSERT INTO categories (id, name)
VALUES
    (1, 'Book'),
    (2, 'Car'),
    (3, 'Gift'),
    (4, 'Phone');

INSERT INTO productcategories (id, product_id, category_id)
VALUES
    (1, 1, 1),
    (2, 1, 3),
    (3, 2, 2);