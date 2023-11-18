from flask import Flask, request, jsonify
import tensorflow as tf
from tensorflow.keras.preprocessing.sequence import pad_sequences
import numpy as np
import pickle

app = Flask(__name__)

# Constants from notebook
VOCAB_SIZE = 10000
MAX_LEN = 250

# Load model
model = tf.keras.models.load_model('sentiment_analysis_model.h5')

# Load the tokenizer
with open('tokenizer.pickle', 'rb') as handle:
    tokenizer = pickle.load(handle)

def encode_text_with_loaded_tokenizer(text, tokenizer):
    tokens = tokenizer.texts_to_sequences([text])
    return pad_sequences(tokens, maxlen=MAX_LEN, padding='post', value=VOCAB_SIZE-1)

@app.route('/analyze', methods=['POST'])
def analyze_sentiment():
    data = request.json
    text = data['text']
    encoded_input = encode_text_with_loaded_tokenizer(text, tokenizer)
    prediction = np.argmax(model.predict(encoded_input))
    sentiment = ['Negative', 'Neutral', 'Positive'][prediction]
    return jsonify({'sentiment': sentiment})

if __name__ == '__main__':
    app.run(debug=True)
