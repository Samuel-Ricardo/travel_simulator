# Travel Simulator

<p align="center"> 
  <a href="https://fullcycle.com.br/" target="_blank">
    <img width="auto" src="https://miro.medium.com/v2/resize:fit:650/1*HkZDCVL7uUpYxOqpR5e_wg.png"/>
  </a> 
</p>

<h4 align="center" >🚀 🟨 Full Cycle Event - 2023 🟨 🚀</h4>

<h4 align="center">
  Application developed during a Programmer Event, the <a style="color: #8a4af3;" href="https://github.com/search?q=imers%C3%A3o%20full%20cycle&type=repositories" target="_blank">Full Cycle Immersion</a> promoted by <a style="color: #8a4af3;" href="https://fullcycle.com.br/" target="_blank">@FullCycleSchool</a>
</h4>

#

<p align="center">
  |&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#project">Overview</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#techs">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#app">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#run-project">Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#author">Author</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

#

<h1 align="center">
  
  <a href="https://github.com/Samuel-Ricardo">
    <img src="https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=GITHUB"/>
  </a>

  <a herf="https://www.instagram.com/samuel_ricardo.ex/">
    <img src='https://img.shields.io/static/v1?label=&message=Samuel.ex&color=black&style=for-the-badge&logo=instagram'/> 
  </a>

  <a herf='https://www.linkedin.com/in/samuel-ricardo/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=LinkedIn'/> 
  </a>

</h1>

<br>

<p id="project"/>

<h2>  | :artificial_satellite: About:  </h2>

<p align="justify">
    This project is a application of a complete Full Cycle project, with amazing technologies like NextJS for front-end, NestJS for back-end, PostgreSQL database, GO Lang for microsservice, Docker & Kubernetes for devops, metrics with Elastic Search & Kibana, Message / Event Driven Architeture with Kafka and more.
</p>

<img src="./readme_files/techs.png"/>

<p align="justify">
  The propurse of this microsservice is only one, simulate and provide travels routes for all applications of this project, for this comunication between apps, i use Kafka that is a amazing solution to handle data traffic using events, super scaleble, garant the safety of every byte of data whithout data loses.
</p>

<p align="justify">
  The Kubernetes manages this applications on Google Cloud Platform, with a amazing peformance and scalability. Kubernetes is a container orchestrator thats ensures resilience to your application using the most advanced strategies for handle your applications
</p>

<p align="justify">
   All goes as data stream having Kafka as intermediator thats ensures no loses keeping the original order of the message even though something crash, when it back to work will run as if nothing has happen, it all is sended to back-end that processes them and send to front-end. All data have backup in Elastic Search thats in addition to besides being a Powerfull Search Engine, it is also a full ecosystem that includes Kibana a full Data Analytics that allow you to use data provideded by Elastic to build powerfull graphics and Analytics
</p>

> <a href="https://samuel-ricardo.github.io/"> <img src="./readme_files/app_preview.png"> </a>

  <br>
  
- This Microsservice is hosted on Google Cloud Platform - [GCP] with Kubernetes 
- Current Version: <b> 1.0.0 </b>

<br>

> <a target="_blank" href="https://cloud.google.com/?hl=pt-br"> <img width="128px" src="https://i.pcmag.com/imagery/reviews/02yVL9f8Jw1atwoG6sgFZDH-7..v1569482492.jpg"/> </a>

<br>

#

<br/>

- Front-End : NextJS | [ [repositories](https://github.com/Samuel-Ricardo/codelivery-site) ]
- Back-End : NestJS | [ [repositories](https://github.com/Samuel-Ricardo/codelivery_api) ]
- microsservice : GO Lang | [ [repositories](https://github.com/Samuel-Ricardo/travel_simulator/tree/main) ]

#

<br>

<h2 id="techs">
  :building_construction: | Technologies and Concepts Studied:
</h2>

> <a href='https://go.dev/'> <img width="64px" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" /> </a>

- Go Lang
- Kafka
- Kubernetes
- Docker
- Elastic Search
- Kibana
- Data Analytics
- Perfomance
- Event Driven Architeture
- Scalability
- Real Time

> Among Others...

<br>

#

<h2 id="app">
  💻 | Application:
</h2>

<img src="./readme_files/scheme.png" />
| Project Scheme

<img src="./readme_files/deploy.png"/>
| Google Cloud Platform

<img src="./readme_files/confluent.png"/>
| Confluent Dashboard - [Kafka]

<img src="./readme_files/elastic_platform.png"/>
| Elastic Platform

<img src="./readme_files/kibana.png"/>
| Kibana Data Analytics

#

<h2 id="run-project"> 
   👨‍💻 | How to use
</h2>

<br>

### Open your Git Terminal and clone this repository

```git
  $ git clone "git@github.com:Samuel-Ricardo/travel_simulator.git"
```

### Make Pull

```git
  $ git pull "git@github.com:Samuel-Ricardo/travel_simulator.git"
```

<br>

This application use `Docker` so you dont need to install and cofigurate anything other than docker on your machine.

> <a target="_blank" href="https://www.docker.com/"> <img width="48px" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-plain-wordmark.svg" /> </a>

<br>

<p align="justify">

First you need to up Apache Kafka and Elastic Search environment, we will use docker for this. so let's to apache folder `$ cd ./.docker/apache_kafka/` and run it.

</p>

```bash

  # Once docker is configured:

  $ cd ./.docker/apache_kafka/
  $ docker-compose up


```

<p align="justify">

Once Apache Kafka, Confluent Dashboard, Zookeper, Elastic Search and Kibana are up, we can start the project :D (to go back for root project folder: `$ cd ../../`)

</p>

```bash

  # After setup docker environment just run this commmand on root project folder:

  $ docker-compose up --build   # For First Time run this command

  $ docker-compose up           # to run project


```

```bash

  #Apps Running on:

  $ Elastic Search: http://localhost:9200

  $ Kibana: http://localhost:5601

  $ [Kafka] Confluent Dashboard: http://localhost:9021

  See more: ./.docker/apache_kafka/docker-compose.yaml

```

<br>

<h2> 
   👨‍💻 | How to run the full project
</h2>

First, you need to setup the Kafka platform with the Travel Microsservice, click on image bellow to setup it

> <a target="_blank" href="https://github.com/Samuel-Ricardo/travel_simulator#------how-to-use"> <img width="128px" src="https://cdn.thenewstack.io/media/2022/05/57bb2a1f-golang.png"/> </a>

Now, you need to setup the NestJS API, click on image bellow to setup it

> <a target="_blank" href="https://github.com/Samuel-Ricardo/codelivery_api#------how-to-use"> <img width="128px" src="https://imgs.search.brave.com/hM_dWTDZ7a9nUL-Xf70pL9QQeiO6v0_q0rHk8NQlNSU/rs:fit:200:200:1/g:ce/aHR0cHM6Ly9uZXN0/anMuY29tL2ltZy9u/ZXN0LW9nLnBuZw"/> </a>

Now, you can setup this NextJS App, click on image bellow to setup it

> <a target="_blank" href="https://github.com/Samuel-Ricardo/codelivery-site#run-project"> <img width="128px" src="https://pbs.twimg.com/card_img/1669374288581853186/RoVDMNTV?format=jpg&name=4096x4096"/> </a>

#

<br>
<br>

#

<h2 id="author">
  :octocat: | Author:  
</h2>

> <a target="_blank" href="https://www.linkedin.com/in/samuel-ricardo/"> <img width="350px" src="https://github.com/Samuel-Ricardo/bolao-da-copa/blob/main/readme_files/IMG_20220904_220148_188.jpg?raw=true"/> <br> <p> <b> - Samuel Ricardo</b> </p></a>

<h1>
  <a herf='https://github.com/Samuel-Ricardo'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=GITHUB'> 
  </a>
  
  <a herf='https://www.instagram.com/samuel_ricardo.ex/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel.ex&color=black&style=for-the-badge&logo=instagram'> 
  </a>
  
  <a herf='https://twitter.com/SamuelR84144340'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=twitter'> 
  </a>
  
   <a herf='https://www.linkedin.com/in/samuel-ricardo/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=LinkedIn'> 
  </a>
</h1>
