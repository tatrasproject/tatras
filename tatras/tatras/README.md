## Summary

FastAPI Python app
Custom Python client for ArgoCD

## Deploy

```bash
git clone https://github.com/tatrasproject/tatras
cd tatras
python virtualenv venv
source venv/bin/activate
pip install tatras/
cd tatras
uvicorn main:app --reload --host 0.0.0.0
```