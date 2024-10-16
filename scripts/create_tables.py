import psycopg2

# Função para ler o arquivo SQL
def read_sql_file(file_path):
    with open(file_path, 'r') as file:
        return file.read()

def create_tables():
    conn = psycopg2.connect(
        dbname='pin_db',
        user='postgres',
        password='123456',
        host='localhost',
        port='5432'
    )
    cursor = conn.cursor()

    ddl_commands = read_sql_file('scripts/content/tables.sql')

    cursor.execute(ddl_commands)

    conn.commit()
    cursor.close()
    conn.close()


if __name__ == "__main__":
    create_tables()