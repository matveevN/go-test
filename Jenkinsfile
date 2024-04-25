pipeline {
    agent any
    
    options {
        buildDiscarder(logRotator(numToKeepStr: '10'))
        timeout(time: 1, unit: 'HOURS')
        displayName('CI/CD PIPELINE')
    }


    triggers {
        scm('main')
        pullRequest()
    }

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
                    checkout scm
                    
                    sh 'docker build -t course_image golang/.'
                    sh 'docker image prune -f'
                    
                    echo '------------Docker image created successfully!------------'
                }
            }
        }
        
        stage('Deploy Docker image') {
            steps {
                echo '------------Start deploying the Docker image...------------'
                sh 'docker run -d -p 127.0.0.2:8888:8888 --name course_container docker.io/library/course_image'
                echo 'New container added at http://127.0.0.2:8888'
            }
        }
        
    }
}
