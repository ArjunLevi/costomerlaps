# Customerlaps TEST

go run main.go

## Using the following curl commands to submit the request

### Requests

```bash
curl -X POST http://localhost:8080/submit -H "Content-Type: application/json" -d '{
    "ev": "contact_form_submitted",
    "et": "form_submit",
    "id": "cl_app_id_001",
    "uid": "cl_app_id_001-uid-001",
    "mid": "cl_app_id_001-uid-001",
    "t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
    "p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
    "l": "en-US",
    "sc": "1920 x 1080",
    "atrk1": "form_varient",
    "atrv1": "red_top",
    "atrt1": "string",
    "atrk2": "ref",
    "atrv2": "XPOWJRICW993LKJD",
    "atrt2": "string",
    "uatrk1": "name",  
    "uatrv1": "iron man",
    "uatrt1": "string",
    "uatrk2": "email",
    "uatrv2": "ironman@avengers.com",
    "uatrt2": "string",
    "uatrk3": "age",
    "uatrv3": "32",
    "uatrt3": "integer"
}'


curl -X POST http://localhost:8080/submit -H "Content-Type: application/json" -d '{
        "ev": "top_cta_clicked",
        "et": "clicked",
        "id": "cl_app_id_001",
        "uid": "cl_app_id_001-uid-001",
        "mid": "cl_app_id_001-uid-001",
        "t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
        "p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
        "l": "en-US",
        "sc": "1920 x 1080",
        "atrk1": "button_text",
        "atrv1": "Free trial",
        "atrt1": "string",
        "atrk2": "color_variation",
        "atrv2": "ESK0023",
        "atrt2": "string",
        "uatrk1": "user_score",
        "uatrv1": "1034",
        "uatrt1": "number",
        "uatrk2": "gender",
        "uatrv2": "m",
        "uatrt2": "string",
        "uatrk3": "tracking_code",
        "uatrv3": "POSERK093",
        "uatrt3": "string"
    }'

curl -X POST http://localhost:8080/submit -H "Content-Type: application/json" -d '{
        "ev": "top_cta_clicked",
        "et": "clicked",
        "id": "cl_app_id_001",
        "uid": "cl_app_id_001-uid-001",
        "mid": "cl_app_id_001-uid-001",
        "t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
        "p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
        "l": "en-US",
        "sc": "1920 x 1080",
        "atrk1": "button_text",
        "atrv1": "Free trial",
        "atrt1": "string",
        "atrk2": "color_variation",
        "atrv2": "ESK0023",
        "atrt2": "string",
        "atrk3": "page_path",
        "atrv3": "/blog/category_one/blog_name.html",
        "atrt3": "string",
        "atrk4": "source",
        "atrv4": "facebook",
        "atrt4": "string",
        "uatrk1": "user_score",
        "uatrv1": "1034",
        "uatrt1": "number",
        "uatrk2": "gender",
        "uatrv2": "m",
        "uatrt2": "string",
        "uatrk3": "tracking_code",
        "uatrv3": "POSERK093",
        "uatrt3": "string",
        "uatrk4": "phone",
        "uatrv4": "9034432423",
        "uatrt4": "number",
        "uatrk5": "coupon_clicked",
        "uatrv5": "true",
        "uatrt5": "boolean",
        "uatrk6": "opt_out",
        "uatrv6": "false",
        "uatrt6": "boolean"
    }'
