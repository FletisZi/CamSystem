CREATE TABLE cameras (
    id BIGSERIAL PRIMARY KEY,

    -- identificador global
    uuid UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),

    -- nome amigável
    name VARCHAR(120) NOT NULL,

    -- identificação da câmera
    ip VARCHAR(45) NOT NULL,
    port INTEGER DEFAULT 554,

    -- rtsp completo opcional
    rtsp_url TEXT,

    -- autenticação
    username VARCHAR(120),
    password_encrypted TEXT,

    -- informações da câmera
    brand VARCHAR(80),
    model VARCHAR(80),

    -- localização
    location VARCHAR(255),

    -- configurações
    fps INTEGER DEFAULT 15,
    resolution_width INTEGER,
    resolution_height INTEGER,

    -- flags
    is_active BOOLEAN DEFAULT true,

    -- auditoria
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


CREATE TABLE video_recordings (

     id BIGSERIAL PRIMARY KEY,

    -- identificador global
    uuid UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),

    camera_id BIGINT REFERENCES cameras(id),

    file_path TEXT NOT NULL,
    file_name TEXT NOT NULL,

    -- started_at TIMESTAMP NOT NULL,
    -- ended_at TIMESTAMP NOT NULL,

    -- duration_seconds INTEGER,

    -- file_size BIGINT,

    checksum_sha256 TEXT,

    event_type VARCHAR(50),

    ticket_id UUID,

    placa_detectada VARCHAR(20),

    -- trigger_source VARCHAR(50),

    -- status VARCHAR(20),

    created_at TIMESTAMP DEFAULT NOW()
);