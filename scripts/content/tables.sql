CREATE TABLE areas (
    id SERIAL PRIMARY KEY,
    area VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE cargos (
    id SERIAL PRIMARY KEY,
    cargo VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE localidades (
    id SERIAL PRIMARY KEY,
    localidade VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE generos (
    id SERIAL PRIMARY KEY,
    genero VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE geracoes (
    id SERIAL PRIMARY KEY,
    geracao VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE funcionarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    email_corporativo VARCHAR(255),
    area_id INTEGER REFERENCES areas(id),
    cargo_id INTEGER REFERENCES cargos(id),
    funcao VARCHAR(255),
    localidade_id INTEGER REFERENCES localidades(id),
    tempo_de_empresa INTEGER,
    genero_id INTEGER REFERENCES generos(id),
    geracao_id INTEGER REFERENCES geracoes(id)
);

CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    funcionario_id INTEGER REFERENCES funcionarios(id),
    feedback TEXT,
    comentarios_feedback TEXT,
    interacao_gestor TEXT,
    comentarios_interacao_gestor TEXT,
    clareza_carreira TEXT,
    comentarios_clareza_carreira TEXT,
    expectativa_permanencia TEXT,
    comentarios_expectativa_permanencia TEXT,
    enps INTEGER,
    aberta_enps TEXT
);
