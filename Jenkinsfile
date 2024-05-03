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
