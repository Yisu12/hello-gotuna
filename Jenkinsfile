pipeline {
    agent any
    timestamp{
        stages {
            stage('Build') {
                steps {
                    sh 'docker-compose build'
                }
            }
            stage('Deploy') {
                steps {
                    sh 'docker-compose up -d'
                }
            }
        }
    }
}
