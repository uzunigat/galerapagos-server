#!groovy

@Library('Jenkins-Audibene-MPL') _

def pipelineConfig = [
  ecrRepoName: 'ta.go-layered-skeletor',
  kubernetesDeploymentName: 'ta.go-layered-skeletor',
  appNamespace: 'teleaudiology',
  ecrRegion: 'eu-central-1',
  newPodLogTailTimeout: '40',
  podRole: 'JenkinsAgentTeleaudiology',
  slackChannel: '#ta-backstage-ci-notifications',
  podContainers: [
    containerTemplate(
      name: 'terraform-pipeline',
      image: '199636132489.dkr.ecr.eu-central-1.amazonaws.com/terraform-pipeline:latest-production',
      ttyEnabled: true,
      privileged: true,
      alwaysPullImage: true,
      command: 'cat'
    )
  ],
  branchToEnvironmentMapping: [
    'develop'  : 'testing',
    'candidate': 'staging',
    'master'   : 'production'
  ],
  environmentToRegionMapping: [
    'testing': ['eu-central-1'],
    'staging': ['eu-central-1'],
    'production': ['eu-central-1']
  ]
]

audPipelineBootstrap(pipelineConfig) {

  stage('AWS CodeArtifact Login'){
    token = sh(returnStdout: true, script: 'aws codeartifact get-authorization-token --region eu-central-1 --domain audibene --domain-owner 199636132489 --query authorizationToken --output text').trim()
  }

  withEnv(["CODEARTIFACT_AUTH_TOKEN=$token"]){
    if (BRANCH_NAME.startsWith('PR-') || BRANCH_NAME == 'develop' || BRANCH_NAME == 'candidate') {
      echo "Running on $BRANCH_NAME"

      stage('Build'){
        sh "docker kill \$(docker ps -q) || docker ps"
        sh "cp .env.ci .env"
        sh "make build"
      }
      stage('Unit Tests'){
        sh "make unit-test"
      }
    }
    if (BRANCH_NAME.startsWith('PR-') || BRANCH_NAME == 'develop') { // run integration test...
      withSecretsInjected(
          [
              "/ta-go-layered-skeletor/${pipelineConfig.environment}/custom_secrets"
          ],
          pipelineConfig.appRegion
      )  {
          stage('Integration Test'){
                sh "cp .env.ci .env"
                sh "make integration-test"
          }
      }
    }
    if (!BRANCH_NAME.startsWith('PR-')) {
      stage('Docker Build & Push'){
        dockerBuildPush {
          ecrUri          = "${pipelineConfig.ecrUri}"
          imageTag        = "${pipelineConfig.gitCommit}"
          environment     = "${pipelineConfig.environment}"
          applyLatestTag  = true
          latestTagName   = "latest-${pipelineConfig.environment}"
          dockerBuildArgs = "--build-arg CODEARTIFACT_AUTH_TOKEN=\${CODEARTIFACT_AUTH_TOKEN}"
        }
        stage('Deploy-' + "${pipelineConfig.environment}") {
          kubernetesDeploy {
            kubernetesDeploymentName = "${pipelineConfig.kubernetesDeploymentName}"
            environment              = "${pipelineConfig.environment}"
            kubernetesClusterName    = "${pipelineConfig.kubernetesClusterName}"
            appNamespace             = "${pipelineConfig.appNamespace}"
            appRegion                = pipelineConfig.environmentToRegionMapping[pipelineConfig.environment]
            ecrUri                   = "${pipelineConfig.ecrUri}"
            imageTag                 = "${pipelineConfig.gitCommit}"
            containerReadyTimeout    = "${pipelineConfig.containerReadyTimeout}"
            newPodLogTailTimeout     = "${pipelineConfig.newPodLogTailTimeout}"
            buildConfig              = pipelineConfig
            useParallelStages        = true
          }
        }
      }
    }
    if (BRANCH_NAME == "develop" || BRANCH_NAME == "candidate") {
      stage('Git Promote'){
        gitPromote {
          codeVersion         = "${pipelineConfig.codeVersion}"
          upstreamBranch      = "${pipelineConfig.gitBranch}" == "develop" ? "candidate" : ("${pipelineConfig.gitBranch}" == "candidate" ? "master" : "")
          abortPreviousBuilds = "${pipelineConfig.gitBranch}" == "develop" ? true : false
          askForApproval      = true
          succeedOnAbort      = true
          buildConfig         = pipelineConfig
        }
      }
    }
  }
}