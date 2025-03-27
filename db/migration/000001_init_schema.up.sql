CREATE TABLE Veterinary_bill (
    id INTEGER PRIMARY KEY,
    username VARCHAR(255),
    price INTEGER,
    created_at TIMESTAMP
);

CREATE TABLE Veterinary_medicine (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    price INTEGER
);

CREATE TABLE Veterinary_medicine_bill (
    bill_id INTEGER,
    medicine_id INTEGER,
    PRIMARY KEY (bill_id, medicine_id),
    FOREIGN KEY (bill_id) REFERENCES Veterinary_bill(id) ON DELETE CASCADE,
    FOREIGN KEY (medicine_id) REFERENCES Veterinary_medicine(id) ON DELETE CASCADE
);
