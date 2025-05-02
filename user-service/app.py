from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_jwt_extended import JWTManager
import os
from dotenv import load_dotenv

load_dotenv()

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = os.getenv('DATABASE_URL', 'postgresql://user:password@localhost/ecommerce')
app.config['JWT_SECRET_KEY'] = os.getenv('JWT_SECRET', 'super-secret-key')
app.config['REDIS_URL'] = os.getenv('REDIS_URL', 'redis://localhost:6379/0')

db = SQLAlchemy(app)
jwt = JWTManager(app)

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), unique=True, nullable=False)
    password = db.Column(db.String(120), nullable=False)
    role = db.Column(db.String(20), default='customer')

@app.route('/')
def health_check():
    return {'status': 'healthy'}

if __name__ == '__main__':
    app.run(debug=True)
