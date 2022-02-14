export const toMermaid = obj => {
    let lines = [
        "flowchart RL"
    ]
    if (obj.spec.resources) {
        obj.spec.resources.forEach(resource => {
            lines.push(`res-${resource.name}["${resource.name}"]`)
            if (resource.sources) {
                resource.sources.forEach(input => {
                    lines.push(`res-${resource.name} --> res-${input.resource}`)
                })
            }
            if (resource.images) {
                resource.images.forEach(input => {
                    lines.push(`res-${resource.name} --> res-${input.resource}`)
                })
            }
            if (resource.configs) {
                resource.configs.forEach(input => {
                    lines.push(`res-${resource.name} --> res-${input.resource}`)
                })
            }
        })
    }
    console.log(lines)
    return lines.join("\n")
}