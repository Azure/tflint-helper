# tflint-helper

A go project containing convenience functions for writing tflint rules.

## modulecontent

The modulecontent package contains functions for extracting information from Terraform modules.
It supports evaluation of all expressions and can be used to extract blocks and attributes according to the supplied inputs.

The blocks and attributes provided by this package can be evaluated using the supplied `*terraform.Evaluator` `EvaluateExpr()`.
Use the `FetchBlocks()` and `FetchAttributes()` functions to extract the blocks and attributes from the module.

You can use the resulting `cty.Value` in the `blockquery` package.

## blockquery

This package queries the `cty.Value` returned by the `modulecontent` package.
The query language based on [gjson](https://github.com/tidwall/gjson), but is not as advanced.

Queries are in dotted string notation, with numeric values used to access list members and a hash symbol for a wildcard.

Use the `Query()` function to return a `cty.Value`.
You can then use one of the comparison functions , e.g. `IsOneOf()` to check the result against a set of expected values.

## rules

These contain template rules for common use cases.

### AzAPI Rule

Use `NewAzApiRule()` to create a rule that checks for specific body properties for a given resource type and API version:

```go
NewAzApiRuleQueryMustExist(
  "ruleName",                                // The rule name
  "https://link-to-rule-docs.com",           // The link to the rule documentation
  "Microsoft.Network/publicIPAddresses",     // The resource type
  "2023-05-01",                              // Minimum API version applicable
  "",                                        // No maximum API version applicable (use latest)
  "properties.sku.name",                     // The query
  blockquery.IsOneOf,                        // The comparison function
  blockquery.NewStringResults("Standard")... // The expected values
)
```
