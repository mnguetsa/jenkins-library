import static com.sap.piper.Prerequisites.checkScript
import groovy.transform.Field

@Field def STEP_NAME = getClass().getName()

void call(Map parameters = [:]) {
    final script = checkScript(this, parameters) ?: this
    String piperGoPath = parameters?.piperGoPath ?: './piper'
    Map cpe = script?.commonPipelineEnvironment?.getCPEMap(script)
    if (cpe == null) {
        return
    }

    def jsonMap = groovy.json.JsonOutput.toJson(cpe)
    script.echo("[MH] PIPER_pipelineEnv: ${jsonMap}")
    if (piperGoPath && jsonMap) {
        withEnv(["PIPER_pipelineEnv=${jsonMap}"]) {
            def output = script.sh(returnStdout: true, script: "${piperGoPath} writePipelineEnv")
            if (parameters?.verbose || true) {
                script.echo("wrote commonPipelineEnvironment: ${output}")
            }
        }
    } else {
        script.echo("can't write pipelineEnvironment: piperGoPath: ${piperGoPath} piperEnvironment ${jsonMap}")
    }
}
