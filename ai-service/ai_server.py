from flask import Flask, request, jsonify
from transformers import AutoTokenizer, AutoModelForCausalLM, AutoModelForSeq2SeqLM, AutoModelForSequenceClassification
import traceback

# Inicializa o Flask
app = Flask(__name__)

# Modelos de geração disponíveis
MODELS = {
    "distilbert": "distilbert-base-uncased",  # Modelo de classificação
    "t5-small": "t5-small",                  # Modelo Seq2Seq
    "flan-t5-base": "google/flan-t5-base",   # Modelo Seq2Seq
    "bloom-560m": "bigscience/bloom-560m"    # Modelo causal
}

# Modelo de tradução
TRANSLATION_MODEL_NAME = "Helsinki-NLP/opus-mt-en-mul"

# Carrega o modelo de tradução e o tokenizer
translation_tokenizer = AutoTokenizer.from_pretrained(TRANSLATION_MODEL_NAME)
translation_model = AutoModelForSeq2SeqLM.from_pretrained(TRANSLATION_MODEL_NAME)

# Cache para modelos de geração carregados
loaded_models = {}

def load_model(model_name):
    """Carrega o modelo de geração e o tokenizer."""
    if model_name not in loaded_models:
        print(f"Loading model: {model_name}")
        tokenizer = AutoTokenizer.from_pretrained(MODELS[model_name])

        # Verifica o tipo de modelo e carrega a classe apropriada
        if "t5" in model_name or "flan" in model_name:
            model = AutoModelForSeq2SeqLM.from_pretrained(MODELS[model_name])
        elif "distilbert" in model_name:
            model = AutoModelForSequenceClassification.from_pretrained(MODELS[model_name])
        else:
            model = AutoModelForCausalLM.from_pretrained(MODELS[model_name])

        loaded_models[model_name] = (tokenizer, model)
    return loaded_models[model_name]

def translate_with_opus_mt(text, target_language):
    """Traduz o texto usando o modelo Helsinki-NLP/opus-mt-en-mul."""
    if target_language == "en":
        return text  # Não traduz se o idioma for inglês
    inputs = translation_tokenizer(f">>{target_language}<< {text}", return_tensors="pt", padding="longest")
    outputs = translation_model.generate(**inputs)
    return translation_tokenizer.decode(outputs[0], skip_special_tokens=True)

@app.route("/generate", methods=["POST"])
def generate():
    """Gera uma mensagem de commit com base no modelo e no prompt fornecidos."""
    try:
        # Dados da solicitação
        data = request.json
        model_name = data.get("model", "t5-small")  # Modelo padrão
        prompt = data.get("prompt", "")
        language = data.get("language", "en")  # Idioma padrão: inglês
        max_length = data.get("max_length", 100)

        # Valida se o modelo solicitado está disponível
        if model_name not in MODELS:
            return jsonify({"error": f"Invalid model: {model_name}"}), 400

        # Carrega o modelo de geração
        tokenizer, model = load_model(model_name)

        if "distilbert" in model_name:
            # Processamento para DistilBERT (retorna uma mensagem fixa, pois não suporta geração de texto)
            return jsonify({
                "commit_message": "DistilBERT não suporta geração de texto.",
                "original_message": "DistilBERT é um modelo de classificação."
            })

        # Tokeniza a entrada corretamente
        inputs = tokenizer(prompt, return_tensors="pt", padding="longest", truncation=True)
        input_ids = inputs["input_ids"]  # Extrai apenas os IDs necessários

        # Gera o texto em inglês
        outputs = model.generate(input_ids=input_ids, max_length=max_length, num_return_sequences=1)
        generated_text = tokenizer.decode(outputs[0], skip_special_tokens=True)

        # Traduz o texto para o idioma solicitado (se diferente de inglês)
        translated_text = translate_with_opus_mt(generated_text, language)

        return jsonify({
            "commit_message": translated_text,
            "original_message": generated_text
        })

    except Exception as e:
        # Captura o erro detalhado para depuração
        error_message = traceback.format_exc()
        print(error_message)  # Loga o erro no terminal
        return jsonify({"error": error_message}), 500

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
