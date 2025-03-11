CREATE TABLE "logs"(
    "id" bigserial PRIMARY KEY,
    "level" varchar(10) NOT NULL,       
    "status_code" int NOT NULL,         
    "method" varchar(10) NOT NULL,      
    "path" text NOT NULL,               
    "latency" interval NOT NULL,        
    "client_ip" varchar(50) NOT NULL,   
    "message" text NOT NULL,        
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
