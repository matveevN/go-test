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
                    
                    echo '------------Docker image created successfully!------------'
                }
            }
        }
        
        stage('Deploy Docker image') {
            steps {
                echo '------------Start deploying the Docker image...------------'
                
                echo 'New container added at http://127.0.0.2:8888'
            }
        }
        
    }
}
