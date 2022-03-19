pipeline {
    agent {
        label 'master'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        ImageName = 'image-name'
        Regist = 'registry.gitlab.com/repo'
        Tag = "${GIT_BRANCH.replaceFirst('/', '-')}-${GIT_COMMIT.substring(0, Math.min(GIT_COMMIT.length(), 7))}"
    }
    stages {
        stage('run-unit-test'){
            steps {
                sh "echo 'run unit test'"
            }
        }
        stage('build-image') {
            steps {
                withDockerRegistry([ credentialsId:'DOCKER_CRED', url:'https://registry.gitlab.com' ]) {
                    sh "docker build -t ${Regist}/${ImageName}:${Tag} ."
                }
            }
        }
        stage('push-image'){
            steps {
                withDockerRegistry([ credentialsId:'DOCKER_CRED', url:'https://registry.gitlab.com' ]) {
                    sh "docker push ${Regist}/${ImageName}:${Tag}"
                    sh "docker rmi ${Regist}/${ImageName}:${Tag}"
                }
            }
        }
        stage('deploy'){
            steps {
                sh "echo 'deploy'"
            }
        }
    }
    post {
        always{
            deleteDir()
            cleanWs externalDelete: 'rm -fr %s'
        }
    }
}
