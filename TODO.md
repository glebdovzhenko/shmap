# General thoughs and plans
* Go is not Python. I can not just take a table, split it into rows and feed rows as context to the templates.
* Templating in general is too powerful for my purposes. I need simple substitutions.

## Templating should go as follows:
1. Select a table
2. **Somehow** rows should be selected
3. We take all columns and selected rows and use that data as substitutions for column names in the template.
