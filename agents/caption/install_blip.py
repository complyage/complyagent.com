# install_blip.py :: Preload BLIP model to reduce cold start

from transformers import BlipProcessor, BlipForConditionalGeneration

print("Downloading BLIP base model...")
BlipProcessor.from_pretrained("Salesforce/blip-image-captioning-base")
BlipForConditionalGeneration.from_pretrained("Salesforce/blip-image-captioning-base")
print("BLIP base model ready.")
