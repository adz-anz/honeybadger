# <img src="./assets/honeybadger.png" alt="drawing" width="48"/> honeybadger

`honeybadger` is a CRUD interface for interacting with the Honeycomb API.

It was inspired by, and takes influence from, the [**`honeymarker`**](https://github.com/honeycombio/honeymarker) interface.

## Installation

```shell
$ go install github.com/adz-anz/honeybadger@latest
$ honeybadger    # (if $GOPATH/bin is in your path.)
```

## Usage

```shell
$ honeybadger -k <your-configuration-key> COMMAND [SUBCOMMAND [subcommand-specific flags]]
```

- `<your-configuration-key>` can be found on `https://ui.honeycomb.io/<team>/environments/<environment>/api_keys`. You can also set it as `HONEYBADGER_CONFIGKEY`
- `COMMAND`/`SUBCOMMAND` see below

## Available Commands

| Implemented        | Command               | Aliases | Description                |
|--------------------|-----------------------|---------|----------------------------|
| :white_check_mark: | `auth`                | `a`     | Manage API Keys            |
| :white_check_mark: | `boards`              | `b`     | Manage Boards              |
| :x:                | `burn_alerts`         | `ba`    | Manage Burn Alerts         |
| :x:                | `columns`             | `c`     | Manage Columns             |
| :white_check_mark: | `datasets`            | `d`     | Manage Datasets            |
| :white_check_mark: | `dataset_defintiions` | `dd`    | Manage Dataset Definitions |
| :x:                | `events`              | `e`     | Manage Events              |
| :white_check_mark: | `markers`             | `m`     | Manage Markers             |
| :white_check_mark: | `marker_settings`     | `ms`    | Manage Marker Settings     |
| :x:                | `queries`             | `q`     | Manage Queries             |
| :x:                | `recipients`          | `r`     | Manage Receipients         |
| :x:                | `slos`                | `s`     | Manage SLOs                |
| :x:                | `triggers`            | `t`     | Manage Triggers            |

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

| Name          | Flag                            | Type     | Description                                                                      | Required           |
|---------------|---------------------------------|----------|----------------------------------------------------------------------------------|--------------------|
| Board Name    | `[-n \| --name] <arg>`          | `string` | The name of the Board.                                                           | :white_check_mark: |
| Description   | `[-d \| --description] <arg>`   | `string` | A description of the Board.                                                      | :x:                |
| Column Layout | `[-c \| --column_layout] <arg>` | `string` | The number of columns to layout on the board. Can be either `multi` or `single`. | :x:                |

#### List All Boards (`boards list`)

> [!NOTE]
> There are no flags currently available for the `boards list` command.

#### Get a Board (`boards get`)

| Name     | Flag                 | Description                            | Required           |
|----------|----------------------|----------------------------------------|--------------------|
| Board ID | `[-i \| --id] <arg>` | The unique identifier (ID) of a Board. | :white_check_mark: |

#### Updating Boards (`boards update`)

| Name          | Flag                            | Type     | Description                                                                      | Required           |
|---------------|---------------------------------|----------|----------------------------------------------------------------------------------|--------------------|
| Board ID      | `[-i \| --id] <arg>`            | `string` | The unique identifier (ID) of a Board.                                           | :white_check_mark: |
| Board Name    | `[-n \| --name] <arg>`          | `string` | The name of the Board.                                                           | :white_check_mark: |
| Description   | `[-d \| --description] <arg>`   | `string` | A description of the Board.                                                      | :x:                |
| Column Layout | `[-c \| --column_layout] <arg>` | `string` | The number of columns to layout on the board. Can be either `multi` or `single`. | :x:                |


#### Deleting Boards (`boards delete`)

| Name     | Flag                 | Type     | Description                            | Required           |
|----------|----------------------|----------|----------------------------------------|--------------------|
| Board ID | `[-i \| --id] <arg>` | `string` | The unique identifier (ID) of a Board. | :white_check_mark: |

#### Add a Query to a Board (`boards add_query`)

| Name                                         | Flag                            | Type     | Description                                                                                                                                                                           | Required           |
|----------------------------------------------|---------------------------------|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| Board ID                                     | `[-i \| --id] <arg>`            | `string` | The unique identifier (ID) of a Board.                                                                                                                                                | :white_check_mark: |
| Query - Captions                             | `[-c \| --caption] <arg>`       | `string` | Descriptive text to contextualize the value of the Query within the Board.                                                                                                            | :x:                |
| Query - Graph Settings - Hide Markers        | `[-H \| --hide_markers]`        | `bool`   | Hide markers on the graph.                                                                                                                                                            | :x:                |
| Query - Graph Settings - Log Scale           | `[-L \| --log_scale]`           | `bool`   | Use a log scale, rather than a linear scale.                                                                                                                                          | :x:                |
| Query - Graph Settings - Omit Missing Values | `[-O \| --omit_missing]`        | `bool`   | Omit missing values from the graph.                                                                                                                                                   | :x:                |
| Query - Graph Settings - Stacked Graphs      | `[-S \| --stacked_graphs]`      | `bool`   | Display groups as stacked colored areas under their line graphs.                                                                                                                      | :x:                |
| Query - Graph Settings - UTC X Axis          | `[-U \| --utx_axis]`            | `bool`   | Displays the X axis in Coordinated Universal Time, the time at 0° longitude.                                                                                                          | :x:                |
| Query - Graph Settings - Overlaid Charts     | `[-V \| --overlaid_charts]`     | `bool`   | Combines any visualized AVG, MIN, MAX, and PERCENTILE clauses into a single chart.                                                                                                    | :x:                |
| Query - Style                                | `[-s \| --style] <arg>`         | `string` | How the query should be displayed on the board. Can be `graph`, `table`, or `combo`.                                                                                                  | :x:                |
| Query - Dataset                              | `[-d \| --dataset] <arg>`       | `string` | The Dataset to Query. Required if using the deprecated query. Note: this field can take either name (`"My Dataset"`) or slug (`"my_dataset"`); the response will always use the name. | :x:                |
| Query - ID                                   | `[-q \| --query_id] <arg>`      | `string` | The ID of a Query object. Cannot be used with query. Query IDs can be retrieved from the UI or from the Query API.                                                                    | :x:                |
| Query - Annotation ID                        | `[-a \| --annotation_id] <arg>` | `string` | The ID of a Query Annotation that provides a name and description for the Query. The Query Annotation must apply to the `query_id` specified.                                         | :x:                |

---

### Managing Datasets (`datasets`)

| Subcommand | Aliases                                 | Description         |
|------------|-----------------------------------------|---------------------|
| `create`   | `add`, `new`                            | Create a Dataset.   |
| `list`     | `ls`                                    | Lists all Datasets. |
| `get`      |                                         | Get a Dataset.      |
| `delete`   | `rm`, `remove`, `del`                   | Delete a Dataset.   |
| `update`   | `up`, `edit`, `modify`, `change`, `set` | Update a Dataset.   |

#### Creating Datasets (`datasets create`)

| Name              | Flag                                | Type     | Description                                        | Required           |
|-------------------|-------------------------------------|----------|----------------------------------------------------|--------------------|
| Name              | `[-n \| --name] <arg>`              | `string` | The name of the dataset.                           | :white_check_mark: |
| Description       | `[-d \| --description] <arg>`       | `string` | A description for the dataset.                     | :x:                |
| Expand JSON Depth | `[-e \| --expand_json_depth] <arg>` | `int`    | The maximum unpacking depth of nested JSON fields. | :x:                |

#### Listing Datasets (`datasets list`)

> [!NOTE]
> There are no flags currently available for the `datasets list` command.

#### Get a Dataset (`datasets get`)

| Name         | Flag                   | Type     | Description       | Required           |
|--------------|------------------------|----------|-------------------|--------------------|
| Dataset Slug | `[-s \| --slug] <arg>` | `string` | The dataset slug. | :white_check_mark: |

#### Deleting Datasets (`datasets delete`)

| Name         | Flag                   | Type     | Description       | Required           |
|--------------|------------------------|----------|-------------------|--------------------|
| Dataset Slug | `[-s \| --slug] <arg>` | `string` | The dataset slug. | :white_check_mark: |

#### Updating Datasets (`datasets update`)

| Name              | Flag                                | Type     | Description                                        | Required           |
|-------------------|-------------------------------------|----------|----------------------------------------------------|--------------------|
| Dataset Slug      | `[-s \| --slug] <arg>`              | `string` | The dataset slug.                                  | :white_check_mark: |
| Description       | `[-d \| --description] <arg>`       | `string` | A description for the dataset.                     | :x:                |
| Expand JSON Depth | `[-e \| --expand_json_depth] <arg>` | `int`    | The maximum unpacking depth of nested JSON fields. | :x:                |

---

### Managing Dataset Definitions (`dataset_definitions`)

| Subcommand | Aliases                                 | Description                  |
|------------|-----------------------------------------|------------------------------|
| `update`   | `up`, `edit`, `modify`, `change`, `set` | Update a Dataset Definition. |
| `get`      | `ls`, `list`                            | Get all Dataset Definitions. |

#### Updating Dataset Definitions (`dataset_definitions update`)

| Name            | Flag                      | Type     | Description                                                                                                                                                                                                  | Required           |
|-----------------|---------------------------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| Dataset Slug    | `--slug <arg>`            | `string` | The dataset slug.                                                                                                                                                                                            | :white_check_mark: |
| Span ID         | `--span-id <arg>`         | `string` | The unique identifier (ID) for each span.                                                                                                                                                                    | :x:                |
| Trace ID        | `--trace-id <arg>`        | `string` | The ID of the trace this span belongs to.                                                                                                                                                                    | :x:                |
| Parent ID       | `--parent-id <arg>`       | `string` | The ID of this span's parent span, the call location the current span was called from.                                                                                                                       | :x:                |
| Name            | `--name <arg>`            | `string` | The name of the function or method where the span was created.                                                                                                                                               | :x:                |
| Service Name    | `--service-name <arg>`    | `string` | The name of the instrumented service.                                                                                                                                                                        | :x:                |
| Duration (ms)   | `--duration-ms <arg>`     | `string` | How much time the span took, in milliseconds.                                                                                                                                                                | :x:                |
| Span Kind       | `--span-kind <arg>`       | `string` | Metadata: Kind - The kind of Span. For example, client or server. The use of this field to identify Span Events and Links is deprecated. Use the `--annotation-type` flag.                                   | :x:                |
| Annotation Type | `--annotation-type <arg>` | `string` | Metadata: Annotation Type - The type of span annotation. For example, span_event or link. This lets Honeycomb visualize this type of event differently in a trace. Do not use this field for other purposes. | :x:                |
| Link Span ID    | `--link-span-id <arg>`    | `string` | Metadata: Link Span ID - Links let you tie traces and spans to one another. The Link Span ID lets you link to a different span (when used with Link Trace ID).                                               | :x:                |
| Link Trace ID   | `--link-trace-id <arg>`   | `string` | Metadata: Link Trace ID - Links let you tie traces and spans to one another. The Link Trace Id lets you link to a different trace or a different span in the same trace (when used with Link Span ID).       | :x:                |
| Error           | `--error <arg>`           | `string` | Use a `string` to indicate an error.                                                                                                                                                                         | :x:                |
| Status          | `--status <arg>`          | `string` | Indicates the success, failure, or other status of a request.                                                                                                                                                | :x:                |
| Route           | `--route <arg>`           | `string` | The HTTP URL or equivalent route processed by the request.                                                                                                                                                   | :x:                |
| User            | `--user <arg>`            | `string` | The user making the request in the system.                                                                                                                                                                   | :x:                |


#### Get All Dataset Definitions (`dataset_definitions get`)

| Name         | Flag           | Type     | Description       | Required           |
|--------------|----------------|----------|-------------------|--------------------|
| Dataset Slug | `--slug <arg>` | `string` | The dataset slug. | :white_check_mark: |

---

### Managing Markers (`markers`)

| Subcommand | Aliases                                 | Description                               |
|------------|-----------------------------------------|-------------------------------------------|
| `create`   | `add`, `new`                            | Create a Marker in the specified dataset. |
| `list`     | `ls`, `get`                             | Lists all Markers for a dataset.          |
| `update`   | `up`, `edit`, `modify`, `change`, `set` | Update a Marker in the specified dataset. |
| `delete`   | `rm`, `remove`, `del`                   | Delete a Marker in the specified dataset. |

> [!IMPORTANT]
> All `markers` subcommands are configured with the `dataset` flag. Although it is **_not required_**, this will default to the `__all__` dataset, which will affect markers across all datasets. This does not mean you can affect change in markers across all datasets - try to think of the `__all__` dataset as a distinct dataset that will be rendered across all datasets.

| Name    | Flag                | Type     | Description                                                                                         | Required |
|---------|---------------------|----------|-----------------------------------------------------------------------------------------------------|----------|
| Dataset | `[-d \| --dataset]` | `string` | The dataset slug or use `__all__` (or omit) for endpoints that support environment-wide operations. | :x:      |

#### Creating Markers (`markers create`)

| Name       | Flag                         | Type     | Description                                                                                                                             | Required |
|------------|------------------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|----------|
| Start Time | `[-s \| --start_time] <arg>` | `int64`  | Indicates the time the Marker should be placed. If missing, defaults to the time the request arrives. Expressed in Unix Time.           | :x:      |
| End Time   | `[-e \| --end_time] <arg>`   | `int64`  | Specifies end time, and allows a Marker to be recorded as representing a time range, such as a 5 minute deploy. Expressed in Unix Time. | :x:      |
| Message    | `[-m \| --msg] <arg>`        | `string` | A message to describe this specific Marker.                                                                                             | :x:      |
| Type       | `[-t \| --type] <arg>`       | `string` | Groups similar Markers. For example, 'deploys'. All Markers of the same type appear with the same color on the graph.                   | :x:      |
| URL        | `[-u \| --url] <arg>`        | `string` | A target for the marker. Clicking the marker text will take you to this URL.                                                            | :x:      |

#### Listing Markers (`markers list`)

> [!NOTE]
> There are no flags currently available for the `markers list` command.

#### Update Markers (`markers update`)

| Name       | Flag                         | Type     | Description                                                                                                                             | Required           |
|------------|------------------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| Marker ID  | `[-i \| --id] <arg>`         | `string` | The unique identifier (ID) of a Marker.                                                                                                 | :white_check_mark: |
| Start Time | `[-s \| --start_time] <arg>` | `int64`  | Indicates the time the Marker should be placed. If missing, defaults to the time the request arrives. Expressed in Unix Time.           | :x:                |
| End Time   | `[-e \| --end_time] <arg>`   | `int64`  | Specifies end time, and allows a Marker to be recorded as representing a time range, such as a 5 minute deploy. Expressed in Unix Time. | :x:                |
| Message    | `[-m \| --msg] <arg>`        | `string` | A message to describe this specific Marker.                                                                                             | :x:                |
| Type       | `[-t \| --type] <arg>`       | `string` | Groups similar Markers. For example, 'deploys'. All Markers of the same type appear with the same color on the graph.                   | :x:                |
| URL        | `[-u \| --url] <arg>`        | `string` | A target for the marker. Clicking the marker text will take you to this URL.                                                            | :x:                |

#### Delete Markers (`markers delete`)

| Name      | Flag                 | Type     | Description                             | Required           |
|-----------|----------------------|----------|-----------------------------------------|--------------------|
| Marker ID | `[-i \| --id] <arg>` | `string` | The unique identifier (ID) of a Marker. | :white_check_mark: |

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

| Name    | Flag                | Type     | Description                                                                                         | Required |
|---------|---------------------|----------|-----------------------------------------------------------------------------------------------------|----------|
| Dataset | `[-d \| --dataset]` | `string` | The dataset slug or use `__all__` (or omit) for endpoints that support environment-wide operations. | :x:      |

#### Creating Markers Settings (`marker_settings create`)

| Name         | Flag                    | Type     | Description                                                                                                            | Required           |
|--------------|-------------------------|----------|------------------------------------------------------------------------------------------------------------------------|--------------------|
| Marker Type  | `[-t \| --type] <arg>`  | `string` | Groups similar Markers. For example, 'deploys'. All Markers of the same type appears with the same color on the graph. | :white_check_mark: |
| Marker Color | `[-c \| --color] <arg>` | `string` | Color to use for display of this marker type. Specified as hexadecimal RGB. For example, `#F96E11`.                    | :white_check_mark: |

#### Get a Marker Setting (`marker_settings get`)

> [!NOTE]
> Despite being described as "Get a Marker Setting", the API endpoint behind this subcommand appears to list all Marker Settings.
> As such, there are no flags currently available for the `marker_settings get` command.

#### Update Marker Settings (`marker_settings update`)

| Name              | Flag                    | Type     | Description                                                                                                            | Required           |
|-------------------|-------------------------|----------|------------------------------------------------------------------------------------------------------------------------|--------------------|
| Marker Setting ID | `[i \| --id] <arg>`     | `string` | The ID of the marker setting to update.                                                                                | :white_check_mark: |
| Marker Type       | `[-t \| --type] <arg>`  | `string` | Groups similar Markers. For example, 'deploys'. All Markers of the same type appears with the same color on the graph. | :white_check_mark: |
| Marker Color      | `[-c \| --color] <arg>` | `string` | Color to use for display of this marker type. Specified as hexadecimal RGB. For example, `#F96E11`.                    | :white_check_mark: |

#### Delete Marker Settings (`marker_settings delete`)

| Name              | Flag                | Type     | Description                             | Required           |
|-------------------|---------------------|----------|-----------------------------------------|--------------------|
| Marker Setting ID | `[i \| --id] <arg>` | `string` | The ID of the marker setting to update. | :white_check_mark: |
