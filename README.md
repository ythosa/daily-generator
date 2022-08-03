# daily-generator

## Description
Generator for daily. Generates by jira task ids message in format:
```markdown
**Что вы делали вчера?**
* [MYPROJECT-321](https://jira.com/MYPROJECT-321) - MYPROJECT-321 issue summary
* [MYPROJECT-322](https://jira.com/MYPROJECT-322) - MYPROJECT-322 issue summary  
**Что вы делали сегодня?**
* [MYPROJECT-323](https://jira.com/MYPROJECT-323) - MYPROJECT-323 issue summary 
**Отлично, есть ли какие-то препятствия?**
Nope
```

## Configuration
### Environment
```env
DG_CONFIGS_FOLDER_PATH=./configs   # configs folder path, "./configs" as default
DG_CONFIG_NAME=config              # config name, "config" as default
```
### Config file (.yaml)
```yaml
jira:
  url: "https://jira.ru"   # Jira base url 
  username: "username"     # Jira username
  password: "password"     # Jira password
  project: "MYPROJECT"     # Jira project

```

## Using
### Steps
1. Run application:
    * ```shell
      make build && ./daily-generator
      ```
   or
   * ```shell
      make run
     ```
2. `>> 🍉 Input yesterday issues:` - here u must input comma separated issue ids 
   * Example: `321, 322`
3. `>> 🍒 Input today issues:` - here u must input comma separated issue ids
   * Example: `323`
4. `>> 🍑 Input problems: Nope` - here u must input description of problems without line breaks
   * Example: `Interviews :((`
### Result
The result will be copied to the clipboard and output to the console. 
Here is result for our examples:
```markdown
**Что вы делали вчера?**
* [MYPROJECT-321](https://jira.com/MYPROJECT-321) - some description from jira 1
* [MYPROJECT-322](https://jira.com/MYPROJECT-322) - some description from jira 2
**Что вы делали сегодня?**
* [MYPROJECT-323](https://jira.com/MYPROJECT-323) - some description from jira 3
**Отлично, есть ли какие-то препятствия?**
Interviews :((
```
