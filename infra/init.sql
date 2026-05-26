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