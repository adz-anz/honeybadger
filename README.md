# honeybadger

`honeybadger` is a CRUD interface for interacting with the Honeycomb API.

It was inspired by, and takes influence from, the [**`honeymarker`**](https://github.com/honeycombio/honeymarker) interface.

## Installation

```shell
$ go install github.com/honeycombio/honeymarker@latest
$ honeymarker    # (if $GOPATH/bin is in your path.)
```

## Usage

```shell
$ ./honeybadger -k <your-configuration-key> COMMAND [SUBCOMMAND [subcommand-specific flags]]
```

- `<your-configuration-key>` can be found on `https://ui.honeycomb.io/<team>/environments/<environment>/api_keys`. You can also set it as `HONEYBADGER_CONFIGKEY`
- `COMMAND`/`SUBCOMMAND` see below

## Available Commands

| Implemented        | Command           | Aliases | Description            |
|--------------------|-------------------|---------|------------------------|
| :white_check_mark: | `auth`            | `a`     | Manage API keys        |
| :white_check_mark: | `boards`          | `b`     | Manage Boards          |
| :x:                | `burn_alerts`     | `ba`    | Manage Burn Alerts     |
| :x:                | `columns`         | `c`     | Manage Columns         |
| :x:                | `datasets`        | `d`     | Manage Datasets        |
| :x:                | `events`          | `e`     | Manage Events          |
| :white_check_mark: | `markers`         | `m`     | Manage Markers         |
| :white_check_mark: | `marker_settings` | `ms`    | Manage Marker Settings |
| :x:                | `queries`         | `q`     | Manage Queries         |
| :x:                | `recipients`      | `r`     | Manage Receipients     |
| :x:                | `slos`            | `s`     | Manage SLOs            |
| :x:                | `triggers`        | `t`     | Manage Triggers        |

---

### Managing API Keys (`auth`)

| Subcommand | Aliases     | Description                                            |
|------------|-------------|--------------------------------------------------------|
| `list`     | `ls`, `get` | List authorizations for the current configuration key. |

#### Listing API Keys (`auth list`)

> [!NOTE]
> There are no flags currently available for the `auth list` command.

---

### Managing Boards (`boards`)

| Subcommand | Aliases                                 | Description                                                      |
|------------|-----------------------------------------|------------------------------------------------------------------|
| `create`   | `add`, `new`                            | Create a Board comprised of one or more Queries.                 |
| `list`     | `ls`                                    | Retrieves a list of all non-secret Boards within an environment. |
| `get`      |                                         | Get a single Board by ID.                                        |
| `update`   | `up`, `edit`, `modify`, `change`, `set` | Update a Board by ID.                                            |
| `delete`   | `rm`, `remove`, `del`                   | Delete a Board by ID.                                            |

#### Creating Boards (`boards create`)

| Name                                         | Flag                            | Description                                                                                                                                                                           | Required           |
|----------------------------------------------|---------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| Board Name                                   | `[-n \| --name] <arg>`          | The name of the Board.                                                                                                                                                                | :white_check_mark: |
| Description                                  | `[-d \| --description] <arg>`   | A description of the Board.                                                                                                                                                           | :x:                |
| Column Layout                                | `[-c \| --column_layout] <arg>` | The number of columns to layout on the board. Can be either `multi` or `single`.                                                                                                      | :x:                |
| Query - Captions                             | `--qc <arg>`                    | Descriptive text to contextualize the value of the Query within the Board.                                                                                                            | :x:                |
| Query - Graph Settings - Hide Markers        | `--qgshm <arg>`                 | Hide markers on the graph.                                                                                                                                                            | :x:                |
| Query - Graph Settings - Log Scale           | `--qgsls <arg>`                 | Use a log scale, rather than a linear scale.                                                                                                                                          | :x:                |
| Query - Graph Settings - Omit Missing Values | `--qgsomv <arg>`                | Omit missing values from the graph.                                                                                                                                                   | :x:                |
| Query - Graph Settings - Stacked Graphs      | `--qgssg <arg>`                 | Display groups as stacked colored areas under their line graphs.                                                                                                                      | :x:                |
| Query - Graph Settings - UTC X Axis          | `--qgsuxa <arg>`                | Displays the X axis in Coordinated Universal Time, the time at 0° longitude.                                                                                                          | :x:                |
| Query - Graph Settings - Overlaid Charts     | `--qgsoc <arg>`                 | Combines any visualized AVG, MIN, MAX, and PERCENTILE clauses into a single chart.                                                                                                    | :x:                |
| Query - Style                                | `--qs <arg>`                    | How the query should be displayed on the board. Can be `graph`, `table`, or `combo`.                                                                                                  | :x:                |
| Query - Dataset                              | `--qd <arg>`                    | The Dataset to Query. Required if using the deprecated query. Note: this field can take either name (\"My Dataset\") or slug (\"my_dataset\"); the response will always use the name. | :x:                |
| Query - ID                                   | `--qid <arg>`                   | The ID of a Query object. Cannot be used with query. Query IDs can be retrieved from the UI or from the Query API.                                                                    | :x:                |
| Query - Annotation ID                        | `--qaid <arg>`                  | The ID of a Query Annotation that provides a name and description for the Query. The Query Annotation must apply to the query_id or query specified.                                  | :x:                |

> [!IMPORTANT]
> All of the `Query` args in the table above are arrays, as the API allows you to associate multiple Queries to a Board.

#### List All Boards (`boards list`)

> [!NOTE]
> There are no flags currently available for the `boards list` command.

#### Get a Board (`boards get`)

| Name     | Flag                 | Description                            | Required           |
|----------|----------------------|----------------------------------------|--------------------|
| Board ID | `[-i \| --id] <arg>` | The unique identifier (ID) of a Board. | :white_check_mark: |

#### Updating Boards (`boards update`)

| Name                                         | Flag                            | Description                                                                                                                                                                           | Required           |
|----------------------------------------------|---------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| Board ID                                     | `[-i \| --id] <arg>`            | The unique identifier (ID) of a Board.                                                                                                                                                | :white_check_mark: |
| Board Name                                   | `[-n \| --name] <arg>`          | The name of the Board.                                                                                                                                                                | :white_check_mark: |
| Description                                  | `[-d \| --description] <arg>`   | A description of the Board.                                                                                                                                                           | :x:                |
| Column Layout                                | `[-c \| --column_layout] <arg>` | The number of columns to layout on the board. Can be either `multi` or `single`.                                                                                                      | :x:                |
| Query - Captions                             | `--qc <arg>`                    | Descriptive text to contextualize the value of the Query within the Board.                                                                                                            | :x:                |
| Query - Graph Settings - Hide Markers        | `--qgshm <arg>`                 | Hide markers on the graph.                                                                                                                                                            | :x:                |
| Query - Graph Settings - Log Scale           | `--qgsls <arg>`                 | Use a log scale, rather than a linear scale.                                                                                                                                          | :x:                |
| Query - Graph Settings - Omit Missing Values | `--qgsomv <arg>`                | Omit missing values from the graph.                                                                                                                                                   | :x:                |
| Query - Graph Settings - Stacked Graphs      | `--qgssg <arg>`                 | Display groups as stacked colored areas under their line graphs.                                                                                                                      | :x:                |
| Query - Graph Settings - UTC X Axis          | `--qgsuxa <arg>`                | Displays the X axis in Coordinated Universal Time, the time at 0° longitude.                                                                                                          | :x:                |
| Query - Graph Settings - Overlaid Charts     | `--qgsoc <arg>`                 | Combines any visualized AVG, MIN, MAX, and PERCENTILE clauses into a single chart.                                                                                                    | :x:                |
| Query - Style                                | `--qs <arg>`                    | How the query should be displayed on the board. Can be `graph`, `table`, or `combo`.                                                                                                  | :x:                |
| Query - Dataset                              | `--qd <arg>`                    | The Dataset to Query. Required if using the deprecated query. Note: this field can take either name (\"My Dataset\") or slug (\"my_dataset\"); the response will always use the name. | :x:                |
| Query - ID                                   | `--qid <arg>`                   | The ID of a Query object. Cannot be used with query. Query IDs can be retrieved from the UI or from the Query API.                                                                    | :x:                |
| Query - Annotation ID                        | `--qaid <arg>`                  | The ID of a Query Annotation that provides a name and description for the Query. The Query Annotation must apply to the query_id or query specified.                                  | :x:                |

> [!IMPORTANT]
> All of the `Query` args in the table above are arrays, as the API allows you to associate multiple Queries to a Board.

#### Deleting Boards (`boards delete`)

| Name     | Flag                 | Description                            | Required           |
|----------|----------------------|----------------------------------------|--------------------|
| Board ID | `[-i \| --id] <arg>` | The unique identifier (ID) of a Board. | :white_check_mark: |

---

### Managing Markers (`markers`)

| Subcommand | Aliases                                 | Description                               |
|------------|-----------------------------------------|-------------------------------------------|
| `create`   | `add`, `new`                            | Create a Marker in the specified dataset. |
| `list`     | `ls`, `get`                             | Lists all Markers for a dataset.          |
| `update`   | `up`, `edit`, `modify`, `change`, `set` | Update a Marker in the specified dataset  |
| `delete`   | `rm`, `remove`, `del`                   | Delete a Marker in the specified dataset. |

> [!IMPORTANT]
> All `markers` subcommands are configured with the `dataset` flag. Although it is **_not required_**, this will default to the `__all__` dataset, which will affect markers across all datasets. This does not mean you can affect change in markers across all datasets - try to think of the `__all__` dataset as a distinct dataset that will be rendered across all datasets.

| Name    | Flag                | Description                                                                                         | Required |
|---------|---------------------|-----------------------------------------------------------------------------------------------------|----------|
| Dataset | `[-d \| --dataset]` | The dataset slug or use `__all__` (or omit) for endpoints that support environment-wide operations. | :x:      |

#### Creating Markers (`markers create`)

| Name       | Flag                         | Description                                                                                                                             | Required |
|------------|------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|----------|
| Start Time | `[-s \| --start_time] <arg>` | Indicates the time the Marker should be placed. If missing, defaults to the time the request arrives. Expressed in Unix Time.           | :x:      |
| End Time   | `[-e \| --end_time] <arg>`   | Specifies end time, and allows a Marker to be recorded as representing a time range, such as a 5 minute deploy. Expressed in Unix Time. | :x:      |
| Message    | `[-m \| --msg] <arg>`        | A message to describe this specific Marker.                                                                                             | :x:      |
| Type       | `[-t \| --type] <arg>`       | Groups similar Markers. For example, 'deploys'. All Markers of the same type appear with the same color on the graph.                   | :x:      |
| URL        | `[-u \| --url] <arg>`        | A target for the marker. Clicking the marker text will take you to this URL.                                                            | :x:      |

#### Listing Markers (`markers list`)

> [!NOTE]
> There are no flags currently available for the `markers list` command.

#### Update Markers (`markers update`)

| Name       | Flag                         | Description                                                                                                                             | Required           |
|------------|------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| Marker ID  | `[-i \| --id] <arg>`         | The unique identifier (ID) of a Marker.                                                                                                 | :white_check_mark: |
| Start Time | `[-s \| --start_time] <arg>` | Indicates the time the Marker should be placed. If missing, defaults to the time the request arrives. Expressed in Unix Time.           | :x:                |
| End Time   | `[-e \| --end_time] <arg>`   | Specifies end time, and allows a Marker to be recorded as representing a time range, such as a 5 minute deploy. Expressed in Unix Time. | :x:                |
| Message    | `[-m \| --msg] <arg>`        | A message to describe this specific Marker.                                                                                             | :x:                |
| Type       | `[-t \| --type] <arg>`       | Groups similar Markers. For example, 'deploys'. All Markers of the same type appear with the same color on the graph.                   | :x:                |
| URL        | `[-u \| --url] <arg>`        | A target for the marker. Clicking the marker text will take you to this URL.                                                            | :x:                |

#### Delete Markers (`markers delete`)

| Name      | Flag                 | Description                             | Required           |
|-----------|----------------------|-----------------------------------------|--------------------|
| Marker ID | `[-i \| --id] <arg>` | The unique identifier (ID) of a Marker. | :white_check_mark: |

---

### Managing Marker Settings (`marker_settings`)

| Subcommand | Aliases                                 | Description                                        |
|------------|-----------------------------------------|----------------------------------------------------|
| `create`   | `add`, `new`                            | Create a Marker Setting in the specified dataset.  |
| `get`      | `ls`, `list`                            | List all Marker Settings in the specified dataset. |
| `update`   | `up`, `edit`, `modify`, `change`, `set` | Update a Marker Setting in the specified dataset   |
| `delete`   | `rm`, `remove`, `del`                   | Delete a Marker Setting in the specified dataset.  |

> [!IMPORTANT]
> All `marker_settings` subcommands are configured with the `dataset` flag. Although it is **_not required_**, this will default to the `__all__` dataset, which will affect marker settings across all datasets. This does not mean you can affect change in marker settings across all datasets - try to think of the `__all__` dataset as a distinct dataset that will be rendered across all datasets.

| Name    | Flag                | Description                                                                                         | Required |
|---------|---------------------|-----------------------------------------------------------------------------------------------------|----------|
| Dataset | `[-d \| --dataset]` | The dataset slug or use `__all__` (or omit) for endpoints that support environment-wide operations. | :x:      |

#### Creating Markers Settings (`marker_settings create`)

| Name         | Flag                    | Description                                                                                                            | Required           |
|--------------|-------------------------|------------------------------------------------------------------------------------------------------------------------|--------------------|
| Marker Type  | `[-t \| --type] <arg>`  | Groups similar Markers. For example, 'deploys'. All Markers of the same type appears with the same color on the graph. | :white_check_mark: |
| Marker Color | `[-c \| --color] <arg>` | Color to use for display of this marker type. Specified as hexadecimal RGB. For example, `#F96E11`.                    | :white_check_mark: |

#### Get a Marker Setting (`marker_settings get`)

> [!NOTE]
> Despite being described as "Get a Marker Setting", the API endpoint behind this subcommand appears to list all Marker Settings.
> As such, there are no flags currently available for the `marker_settings get` command.

#### Update Marker Settings (`marker_settings update`)

| Name              | Flag                    | Description                                                                                                            | Required           |
|-------------------|-------------------------|------------------------------------------------------------------------------------------------------------------------|--------------------|
| Marker Setting ID | `[i \| --id] <arg>`     | The ID of the marker setting to update.                                                                                | :white_check_mark: |
| Marker Type       | `[-t \| --type] <arg>`  | Groups similar Markers. For example, 'deploys'. All Markers of the same type appears with the same color on the graph. | :white_check_mark: |
| Marker Color      | `[-c \| --color] <arg>` | Color to use for display of this marker type. Specified as hexadecimal RGB. For example, `#F96E11`.                    | :white_check_mark: |

#### Delete Marker Settings (`marker_settings delete`)

| Name              | Flag                | Description                             | Required           |
|-------------------|---------------------|-----------------------------------------|--------------------|
| Marker Setting ID | `[i \| --id] <arg>` | The ID of the marker setting to update. | :white_check_mark: |
