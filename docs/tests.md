# Tests and TDD

It is a goal of the project to be developed using TDD. Let's see if we can keep this promise.

I believe that if we follow TDD practices, we will achieve a high percentage of test coverage, but we do not have a minimum coverage requirement for acceptance.

# Test Framework

**For the back-end**, after testing a couple of test frameworks/suites/packages, we've decided to use [Ginkgo](https://onsi.github.io/ginkgo/) as the framework.

**For the front-end**, we have nothing decided yet.

# Types of tests

We will try to follow the mental model of the "Test Pyramid" (https://semaphore.io/blog/testing-pyramid).

## Unit Test

### Back-end

Following the Go tradition, unit tests will be placed in the same folder as the unit being tested, with the "_test" postfix.

## Integration Test

### Back-end

For the integration tests, we will use an "integration-test" folder at the root level of the back-end folder.

We haven’t decided yet if there will be a `.feature` file describing the integrations.

### End-to-End Test (E2E)

There is a folder "e2e-tests" in the root of the project with scenarios described in `.feature` files using Gherkin language to guide both back-end and front-end development.
