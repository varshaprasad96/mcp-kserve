from sklearn.datasets import load_iris
from sklearn.ensemble import RandomForestClassifier
import joblib
import os

# Create models directory if it doesn't exist
os.makedirs('models', exist_ok=True)

# Load the iris dataset
iris = load_iris()
X, y = iris.data, iris.target

# Train a simple random forest classifier
model = RandomForestClassifier(n_estimators=100)
model.fit(X, y)

# Save the model
model_path = 'models/iris_model.joblib'
joblib.dump(model, model_path)
print(f"Model saved to {model_path}") 