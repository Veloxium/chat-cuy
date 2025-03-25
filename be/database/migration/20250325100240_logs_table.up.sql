CREATE TABLE logs (
    id BIGSERIAL PRIMARY KEY,
    level VARCHAR(10) NOT NULL,       
    status_code INT NOT NULL,         
    method VARCHAR(10) NOT NULL,      
    path TEXT NOT NULL,               
    latency INTERVAL NOT NULL,        
    client_ip VARCHAR(50) NOT NULL,   
    message VARCHAR(100) NOT NULL,        
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
