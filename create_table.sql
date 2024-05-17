CREATE TABLE contact (
    contact_id INTEGER primary key autoincrement,
    name TEXT,
    first_name TEXT,
    last_name TEXT, 
    gender_id INTEGER,
    dob DATE,
    email TEXT,
    phone TEXT,
    address TEXT,
    photo_path TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT
);

CREATE TABLE bank_transaction (
    bank_transaction_id INTEGER primary key autoincrement,
    amount INTEGER,
    bank_account_id TEXT,
    status INTEGER
);