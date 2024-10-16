import create_tables  # Importa o script para criar as tabelas
import import_data  # Importa o script para importar os dados

def main():
    # Executa a função que cria as tabelas
    create_tables.create_tables() 

    # Executa a função que importa os dados
    import_data.import_data()  

if __name__ == "__main__":
    main()
