from sqlalchemy.engine.url import URL
from sqlalchemy import create_engine
from sqlalchemy import inspect

l_val_t=1000

def test(tt):
    postgres_db = {'drivername': 'postgres',
                   'database': 'rails_tpl',
                   'username': 'figo',
                   'password': '123456',
                   'host': 'localhost',
                   'port': 5432}
    engine = create_engine(URL(**postgres_db).__to_string__())
    inspector = inspect(engine)
    print(inspector.get_table_names())
    print(inspector.get_columns('users'))
    for _t in inspector.get_table_names():
        print(inspector.get_columns(_t))
    print(tt)
    return "hello"
