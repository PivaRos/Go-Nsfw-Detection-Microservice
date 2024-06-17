# Go-Nsfw-Detection-Microservice

#### A microservice built with Go and kafka for detecting NSFW content, including nudity and blood, in images. The service uses the nudenet library for detection and returns results with confidence scores ranging from 0 to 1.

## Features

- Detects nudity and blood in images.
- Returns confidence scores (0 to 1) for detected content.
- Integrates with Python for image processing and detection.
- Handles base64-encoded image inputs.

## Development

### Make sure that you have Zookeeper and Kafka Running by:

#### Start local Zookeeper and Kafka:

```
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
```

and

```
kafka-server-start /usr/local/etc/kafka/server.properties
```

#### Or have them start automatically whenever you boot your system.

### Setup PostgreSQL locally

using docker:

```
docker pull postgres:latest

docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=myuser -e POSTGRES_DB=mydatabase -p 5432:5432 -d postgres:latest
```

create table:

```
CREATE TABLE image_classifications (
    image_guid UUID PRIMARY KEY,
    classification_result JSONB
);
```

### Start the application

```
cd module

go run ./
```

if setup correctly the application should be running and ready to get traffic !

## Create events:

### image_upload

create kafka event in the following structure:

```
Topic:"image_upload",
value:{
    imageGuid:uuid
    base64:string
}
```

the output should be for the following type:

```
Topic:"image_classified",
Value:{
    imageGuid:uuid
    classificationResult:{
        class:string,
        score:float
        box:int[]
    }[]
}
```

the output is also being saved to the image_classifications table in the same format for future web api access.

#### Example :

input:

```
topic: image_upload
Value:{
    "imageGuid":"e708b569-3847-4b57-a1c7-0644e60f5e93",
    "base64":"SkZJRgABAQAAAQABAAD...."
}
```

output:

```
Topic:image_classified
Value:{"imageGuid":"e708b569-3847-4b57-a1c7-0644e60f5e93","classificationResult":[{"class":"MALE_BREAST_EXPOSED","score":0.82666015625,"box":[132,220,74,58]},{"class":"MALE_BREAST_EXPOSED","score":0.71337890625,"box":[332,220,75,60]},{"class":"ARMPITS_EXPOSED","score":0.69189453125,"box":[93,208,43,43]},{"class":"BELLY_EXPOSED","score":0.6865234375,"box":[310,300,93,103]},{"class":"BELLY_EXPOSED","score":0.685546875,"box":[151,296,75,105]},{"class":"FACE_MALE","score":0.63916015625,"box":[163,67,69,68]},{"class":"FEMALE_GENITALIA_COVERED","score":0.61181640625,"box":[317,433,50,59]},{"class":"FACE_MALE","score":0.60205078125,"box":[330,87,71,67]},{"class":"ARMPITS_EXPOSED","score":0.5908203125,"box":[404,189,54,47]},{"class":"FEMALE_GENITALIA_COVERED","score":0.377685546875,"box":[172,433,46,56]},{"class":"MALE_BREAST_EXPOSED","score":0.2607421875,"box":[202,229,36,53]}]}
```
