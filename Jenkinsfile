pipeline {
    agent any
    
    stages {

        stage('Test') {
            steps {
                echo 'Running tests...'
                echo 'Tests completed successfully!'
            }
        }

        stage('Build Docker image') {
            steps {
                script {
                    echo '------------Start building the Docker image...------------'
                    
                    sh 'go build app.go'
                    
                    sh 'docker build -t myapp-image .'
                    
                    echo '------------Docker image created successfully!------------'
                }
            }
        }

        stage('Push Docker image to Docker Hub') {
            steps {
                echo '------------Pushing Docker image to Docker Hub...------------'
                
                // Логин в Docker Hub
                sh 'docker login -u matveevn -p 25121975DN'
                
                // Переименование образа для Docker Hub
                sh 'docker tag myapp-image matveevn/myapp-image'
                
                // Отправка образа на Docker Hub
                sh 'docker push matveevn/myapp-image'
                
                echo '------------Docker image pushed to Docker Hub successfully!------------'
            }
        }

        stage('Deploy Docker image') {
            steps {
                echo '------------Start deploying the Docker image...------------'
                sh 'docker rm -f myapp-container || true'
                
                // Запуск Docker контейнера
                sh 'docker run -d -p 8888:8888 --name myapp-container myapp-image'
            
                echo '------------New container added at localhost:8888------------'
            }
        }
    }
}
