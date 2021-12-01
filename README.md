# amazon-url-scrapes
There are two services:
1. To Scrape the product details from amazon url
2. To Store the retrieved product data in a file with createdAt timestamp. It has a api to show the data of file on web

https://user-images.githubusercontent.com/22277084/144312287-308378bf-895e-4dc3-bd4f-57e880e35f61.mp4

The images of both the services are pushed to [DockerHub](https://hub.docker.com/repository/docker/shilpi57/amazon-scrape).

The two services can be run using below command
`$ docker-compose up`

### There are three Rest API's exposed
1. #### http://localhost:8080/product/amazon - scrapes details for given amazon url
    Method: POST
    
    ```
    RequestBody:
      {
    "producturl":"https://www.amazon.in/OnePlus-Smart-Band-Saturation-Compatible/dp/B07XY9BZPM"
        }

    ResponseBody:
        {
    "url": "https://www.amazon.in/OnePlus-Smart-Band-Saturation-Compatible/dp/B07XY9BZPM",
    "product": {
        "name": "Amazon.in: Buy OnePlus Smart Band: 13 Exercise Modes, Blood Oxygen Saturation (SpO2), Heart Rate & Sleep Tracking, 5ATM+Water & Dust Resistant( Android & iOS Compatible) Online at Low Prices in India | OnePlus Reviews & Ratings",
        "imageURL": [
            "https://m.media-amazon.com/images/I/410vmwQSZaL._SS40_.jpg",
            "https://m.media-amazon.com/images/I/51NI8ZEVRfL._SS40_BG85,85,85_BR-120_PKdp-play-icon-overlay__.jpg",
            "https://m.media-amazon.com/images/I/31Ip5bIhL8L._SS40_.jpg",
            "https://m.media-amazon.com/images/I/41c1XANwDpS._SS40_.jpg",
            "https://m.media-amazon.com/images/I/41FWKZaZTaS._SS40_.jpg",
            "https://m.media-amazon.com/images/I/41rOES82gsS._SS40_.jpg",
            "https://m.media-amazon.com/images/I/416jb9ZVasS._SS40_.jpg",
            "https://images-eu.ssl-images-amazon.com/images/I/410vmwQSZaL._SX300_SY300_QL70_ML2_.jpg"
        ],
        "description": "Removable main tracker design allows for effortless transition between dynamic dual-color strap combos. Battery life : Up to 14 days, Battery capacity: 100 mAh\n\n\nOn-demand daytime spot checks and continuous sleep monitoring of blood oxygen saturation (Sp02) quickly and accurately highlight potential health issues.\n\n\nAccess key mobile features directly from your wrist â includes music, camera shutter controls, call - message notifications and many more.\n\n\nAsides from OTA Software Updates, the OnePlus Health App analyzes health data, provides insights and advice on your personal health. OnePlus Health App for iOS platform now available.\n\n\n5ATM and IP68 certified, the band is dust and water resistant up to 50 meters for 10 minutes.",
        "price": "₹2,484.00",
        "totalReviews": "14"
        }
        }
        ```
----
2. #### http://storage-service:8081/save/product/amazon - saves the retreived product detail in file and exposes throug `/data` api
    Method: POST
----
3. #### http://localhost:8081/data - displays the product data stored in file from previous api
    Method: GET
        
    ```
    ResponseBody:
     [   {
        "createdAt": "2021-12-02 01:52:55.7499068 +0530 IST",
        "url": "https://www.amazon.in/OnePlus-Smart-Band-Saturation-Compatible/dp/B07XY9BZPM",
        "product": {
            "name": "Amazon.in: Buy OnePlus Smart Band: 13 Exercise Modes, Blood Oxygen Saturation (SpO2), Heart Rate & Sleep Tracking, 5ATM+Water & Dust Resistant( Android & iOS Compatible) Online at Low Prices in India | OnePlus Reviews & Ratings",
            "imageURL": [
                "https://m.media-amazon.com/images/I/410vmwQSZaL._SS40_.jpg",
                "https://m.media-amazon.com/images/I/51NI8ZEVRfL._SS40_BG85,85,85_BR-120_PKdp-play-icon-overlay__.jpg",
                "https://m.media-amazon.com/images/I/31Ip5bIhL8L._SS40_.jpg",
                "https://m.media-amazon.com/images/I/41c1XANwDpS._SS40_.jpg",
                "https://m.media-amazon.com/images/I/41FWKZaZTaS._SS40_.jpg",
                "https://m.media-amazon.com/images/I/41rOES82gsS._SS40_.jpg",
                "https://m.media-amazon.com/images/I/416jb9ZVasS._SS40_.jpg",
                "https://images-eu.ssl-images-amazon.com/images/I/410vmwQSZaL._SX300_SY300_QL70_ML2_.jpg"
            ],
            "description": "Removable main tracker design allows for effortless transition between dynamic dual-color strap combos. Battery life : Up to 14 days, Battery capacity: 100 mAh\n\n\nOn-demand daytime spot checks and continuous sleep monitoring of blood oxygen saturation (Sp02) quickly and accurately highlight potential health issues.\n\n\nAccess key mobile features directly from your wrist â includes music, camera shutter controls, call - message notifications and many more.\n\n\nAsides from OTA Software Updates, the OnePlus Health App analyzes health data, provides insights and advice on your personal health. OnePlus Health App for iOS platform now available.\n\n\n5ATM and IP68 certified, the band is dust and water resistant up to 50 meters for 10 minutes.",
            "price": "₹2,484.00",
            "totalReviews": "14"
            }
        }
    ]
    ```

----
    



