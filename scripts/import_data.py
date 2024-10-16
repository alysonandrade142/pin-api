import pandas as pd
import psycopg2
from psycopg2 import sql

conn = psycopg2.connect(
    dbname='pin_db',
    user='postgres',
    password='123456',
    host='localhost',
    port='5432'
)
cursor = conn.cursor()

def insert_data(table, columns, values):
    query = sql.SQL("INSERT INTO {} ({}) VALUES ({}) ON CONFLICT DO NOTHING").format(
        sql.Identifier(table),
        sql.SQL(', ').join(map(sql.Identifier, columns)),
        sql.SQL(', ').join(sql.Placeholder() * len(values))
    )
    cursor.execute(query, values)

def process_time_of_company(value):
    if isinstance(value, str):
        parts = value.split()
        for part in parts:
            if part.isdigit():
                return int(part)
    return None


def import_data():

    # Ler o arquivo CSV
    df = pd.read_csv('scripts/content/data.csv', delimiter=';')

    for area in df['area'].unique():
        insert_data('areas', ['area'], (area,))

    for cargo in df['cargo'].unique():
        insert_data('cargos', ['cargo'], (cargo,))

    for localidade in df['localidade'].unique():
        insert_data('localidades', ['localidade'], (localidade,))

    for genero in df['genero'].unique():
        insert_data('generos', ['genero'], (genero,))

    for geracao in df['geracao'].unique():
        insert_data('geracoes', ['geracao'], (geracao,))

    for _, row in df.iterrows():
        cursor.execute("SELECT id FROM areas WHERE area = %s", (row['area'],))
        area_id = cursor.fetchone()[0]

        cursor.execute("SELECT id FROM cargos WHERE cargo = %s", (row['cargo'],))
        cargo_id = cursor.fetchone()[0]

        cursor.execute("SELECT id FROM localidades WHERE localidade = %s", (row['localidade'],))
        localidade_id = cursor.fetchone()[0]

        cursor.execute("SELECT id FROM generos WHERE genero = %s", (row['genero'],))
        genero_id = cursor.fetchone()[0]

        cursor.execute("SELECT id FROM geracoes WHERE geracao = %s", (row['geracao'],))
        geracao_id = cursor.fetchone()[0]

        tempo_de_empresa = process_time_of_company(row['tempo_de_empresa'])

        insert_data('funcionarios', ['nome', 'email', 'email_corporativo', 'area_id', 'cargo_id', 
                                    'funcao', 'localidade_id', 'tempo_de_empresa', 
                                    'genero_id', 'geracao_id'], 
                    (row['nome'], row['email'], row['email_corporativo'], area_id, 
                    cargo_id, row['funcao'], localidade_id, tempo_de_empresa, 
                    genero_id, geracao_id))

    for _, row in df.iterrows():
        cursor.execute("SELECT id FROM funcionarios WHERE nome = %s AND email = %s", 
                    (row['nome'], row['email']))
        funcionario_id = cursor.fetchone()[0]

        insert_data('feedback', ['funcionario_id', 'feedback', 'comentarios_feedback', 
                                'interacao_gestor', 'comentarios_interacao_gestor', 
                                'clareza_carreira', 'comentarios_clareza_carreira', 
                                'expectativa_permanencia', 'comentarios_expectativa_permanencia', 
                                'enps', 'aberta_enps'], 
                    (funcionario_id, row['Feedback'], 
                    row['Comentários - Feedback'], 
                    row['Interação com Gestor'], 
                    row['Comentários - Interação com Gestor'], 
                    row['Clareza sobre Possibilidades de Carreira'], 
                    row['Comentários - Clareza sobre Possibilidades de Carreira'], 
                    row['Expectativa de Permanência'], 
                    row['Comentários - Expectativa de Permanência'], 
                    row['eNPS'], 
                    row['[Aberta] eNPS']))

    conn.commit()
    cursor.close()
    conn.close()


if __name__ == "__main__":
    import_data()