import psycopg2

# Função para ler o arquivo SQL
def read_sql_file(file_path):
    with open(file_path, 'r') as file:
        return file.read()

def create_tables():
    # Conexão com o banco de dados
    conn = psycopg2.connect(
        dbname='api_db',
        user='postgres',
        password='123456',
        host='localhost',
        port='5432'
    )
    cursor = conn.cursor()

    # Lendo o DDL do arquivo
    ddl_commands = read_sql_file('content/tables.sql')

    # Executar os comandos DDL
    cursor.execute(ddl_commands)

    # Commit e fechar a conexão
    conn.commit()
    cursor.close()
    conn.close()


if __name__ == "__main__":
    create_tables()