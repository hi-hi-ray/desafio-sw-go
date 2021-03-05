package errors

const MissingConfigVariable = "The following variables in config toml file is missing: "
const ProblemAtReadingFile = "The following problem occurred when trying to read the toml file: "
const SwapiConsumerError = "The following error occurred when trying to request the api: "
const ReadBodyError = "The following error occurred when trying to read the body: "
const MarshallBodyError = "The following error occurred when trying to parse the component: "
const SwapiInternalError = "An unmapped error happened."
const SwapiNotFoundError = "The planet was not found."
const PlanetExist = "This planet already exits, check him using his id. You can use the method GET to find him."
const PlanetMandatoryFields = "Check the fields name, terrain and climate they are mandatory"
const PlanetDoesNotExist = "This planet doesn't exist, check his id. You can use the method GET to find him."
