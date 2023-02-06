pipeline {
    agent any
    options {
        timestamps()
	    ansiColor('xterm')
        }
        stages {
            stage('Build') {
                steps {
			        echo '\033[42m\033[97mConstruyendo...\033[0m'
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

