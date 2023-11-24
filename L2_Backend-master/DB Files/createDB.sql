CREATE DATABASE book_library2;

CREATE TABLE IF NOT EXISTS libraries(
   Id             SERIAL PRIMARY KEY     NOT NULL,
   Name           TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS users(
   Id             SERIAL PRIMARY KEY     NOT NULL,
   Name           TEXT,
   Email          TEXT UNIQUE,
   Contact_Number  TEXT,
   Role           TEXT CHECK (role = 'admin' OR role = 'reader' OR role = 'owner'),
   Lib_Id          INT references libraries(Id)
   Password        TEXT NOT NULL,
);

CREATE TABLE IF NOT EXISTS book_inventories(
    ISBN              SERIAL PRIMARY KEY     NOT NULL, 
    Lib_Id             INT NOT NULL references libraries(Id), 
    Title             TEXT, 
    Authors           TEXT, 
    Publisher         TEXT, 
    Version           TEXT, 
    Total_Copies       INT, 
    Available_Copies   INT
);

CREATE TABLE IF NOT EXISTS request_events(
    Req_ID              SERIAL PRIMARY KEY     NOT NULL, 
    Book_ID             INT NOT NULL references book_inventories(ISBN), 
    Reader_ID           TEXT NOT NULL, 
    Request_Date        DATE, 
    Approval_Date       DATE, 
    Approver_ID         TEXT,
    Request_Type        TEXT
);

CREATE TABLE IF NOT EXISTS issue_registeries(
    Issue_ID               SERIAL PRIMARY KEY   NOT NULL,
    ISBN                  INT NOT NULL references book_inventories(ISBN), 
    Reader_ID              TEXT NOT NULL, 
    Issue_Approver_ID       TEXT NOT NULL, 
    Issue_Status           TEXT NOT NULL, 
    Issue_Date             DATE NOT NULL,
    Expected_Return_Date    DATE, 
    Return_Date            DATE, 
    Return_Approver_ID      TEXT
);
