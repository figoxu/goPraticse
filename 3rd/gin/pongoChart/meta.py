from sqlalchemy.engine.url import URL
from sqlalchemy import create_engine
from sqlalchemy import inspect


def get_table_names(drivername,database,username,password,host,port):
    postgres_db = {'drivername': drivername,'database': database,'username': username,'password': password,'host': host,'port': int(port)}
    engine = create_engine(URL(**postgres_db).__to_string__())
    inspector = inspect(engine)
    return ','.join(inspector.get_table_names())

def get_columns(tablename,drivername,database,username,password,host,port):
    postgres_db = {'drivername': drivername,
                   'database': database,
                   'username': username,
                   'password': password,
                   'host': host,
                   'port': int(port)}
    engine = create_engine(URL(**postgres_db).__to_string__())
    inspector = inspect(engine)
    return str(inspector.get_columns(tablename))


def test(drivername,database,username,password,host,port):
    return drivername+','+database+','+username+','+password+','+host+','+port

