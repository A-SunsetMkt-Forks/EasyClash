<template>
  <el-card
    class="box-card"
    shadow="never"
  >
    <template #header>
      <div class="card-header">
        <el-row justify="space-around">
          <el-col :span="18">
            <span>{{ props.name }}</span>
          </el-col>
          <el-col
            :span="6"
            style="text-align: right"
          >
            <slot name="button" />
            <el-button
              v-if="props.importBtn"
              class="button"
              type="text"
              @click="dialog = true"
            >导入</el-button>
            <el-button
              class="button"
              type="text"
              @click="addElement"
            >添加</el-button>
          </el-col>
        </el-row>
      </div>
    </template>
    <el-table
      :data="list"
      stripe
      style="width: 100%"
    >
      <el-table-column
        v-for="(item, index) in props.header"
        :key="index + '-' + item.field"
        :prop="item.field"
        :label="item.label || item.field"
      >
        <template #default="scope">
          <Render :content="item.render ? item.render(scope, cb) : scope.row[item.field]" />
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            v-if="scope.row.canEdit"
            size="small"
            @click="handleEdit(scope.$index)"
          >Edit</el-button>
          <el-button
            v-if="scope.row.canDel"
            size="small"
            type="danger"
            @click="handleDelete(scope.$index)"
          >Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <el-dialog
    v-model="formModal"
    title="Tips"
    width="80%"
    :before-close="closeForm"
    destroy-on-close
  >
    <span data-wails-no-drag>
      <VForm
        v-model="value"
        :form-items="formItems"
        :mod="formMod"
        @submit="onSubmit"
        @reset="onReset"
      />
    </span>
  </el-dialog>

  <el-dialog
    v-model="dialog"
    title="Tips"
    width="80%"
    :before-close="handleClose"
    destroy-on-close
  >
    <v-editor
      v-model="source"
      vs-path="/assets/monaco-editor/vs"
    />
    <el-button @click="importConf">确定</el-button>
  </el-dialog>
</template>

<script setup lang="tsx">
import { ElMessage } from 'element-plus'
import { ref, defineProps, PropType, onMounted, computed, watch } from 'vue'
import { VForm } from '@okiss/vform'
import { isFunc, isArray } from '@okiss/utils'
import yaml from 'js-yaml'
import { VEditor } from '@okiss/vform'
import { useAppStore } from '../store/app'
import { intersection } from 'lodash'

const appStore = useAppStore()

const dialog = ref(false)
const handleClose = () => {
  dialog.value = false
}
const source = ref()
const importConf = () => {
  yaml.loadAll(source.value, (data: any) => {
    dialog.value = false
    listSafePush(props.importHandle(data)) && saveElements()
  })
}

const Render = (ps: { content: () => any }) => {
  if (isFunc(ps.content)) {
    return ps.content()
  }
  let show = ps.content

  if (isArray(show)) {
    show = (show).join(', ')
  }

  return <div>{show}</div>
}

interface SaveFun {
  (data: List): Promise<boolean>
}

interface DelFun {
  (data: Record<string, any>): Promise<boolean>
}

interface GetFun {
  (): Promise<List>
}

interface ImportFun {
  (data: List): List
}

const props = defineProps({
  name: {
    type: String,
    default: ''
  },
  listHandle: {
    type: Function as PropType<GetFun>,
    default: () => {
      return new Promise((resolve) => {
        resolve([])
      })
    }
  },
  saveHandle: {
    type: Function as PropType<SaveFun>,
    default: () => {
      return new Promise((resolve) => {
        resolve(true)
      })
    }
  },
  canDel: {
    type: Function as PropType<DelFun>,
    default: () => {
      return new Promise((resolve) => {
        resolve(true)
      })
    }
  },
  importHandle: {
    type: Function as PropType<ImportFun>,
    default: () => {
      return new Promise((resolve) => {
        resolve([])
      })
    }
  },
  form: {
    type: [Array, Function],
    default: () => []
  },
  header: {
    type: Array,
    default: () => []
  },
  importBtn: {
    type: Boolean,
    default: true
  },
  insertType: {
    type: String,
    default: 'push'
  }
})

type List = Array<Record<string, string>>
type MapOb = Record<string, any>

const list = ref<List>([])

const checkPre = (o: List): string => {
  const names: Array<string> = []
  o.forEach(e => {
    if (names.indexOf(e.name) === -1) {
      names.push(e.name)
    }
  })
  if (o.length !== names.length) {
    return '导入的数据中名称有重复'
  }
  const names2: Array<string> = []
  list.value.forEach(e => {
    names2.push(e.name)
  })

  if (intersection(names2, names).length !== 0) {
    return '该name已存在'
  }
  return ''
}

const listSafePush = (o: List): boolean => {
  const err = checkPre(o)
  if (err !== '') {
    ElMessage.error(err)
    return false
  }
  list.value = list.value.concat(o)
  return true
}
function reloadList() {
  props.listHandle().then((res: List) => {
    list.value = res || []
  })
}

onMounted(() => {
  reloadList()
})

const value = ref({})
const formItems = computed(() => {
  if (isFunc(props.form)) {
    return (props.form as CallableFunction)()
  }
  return props.form
})

const onSubmit = (formData: MapOb) => {
  const d = Object.assign({}, { ...formData })
  let res = false
  if (editIndex.value > -1) {
    list.value[editIndex.value] = d
    res = true
  } else {
    if (props.insertType === 'push') {
      res = listSafePush([d])
    } else {
      list.value.unshift(d)
      res = true
    }
  }
  closeForm()
  res && saveElements()
}

const onReset = () => {
  formModal.value = false
}

const editIndex = ref(-1)
const handleEdit = (index: number) => {
  formMod.value = 'edit'
  editIndex.value = index
  openForm()
}
const handleDelete = (index: number) => {
  props.canDel(list.value[index]).then(res => {
    if (res) {
      list.value.splice(index, 1)
      saveElements()
    }
  })
}

const formModal = ref(false)
const formMod = ref('create')
const openForm = () => {
  formModal.value = true
  if (editIndex.value !== -1) {
    value.value = { ...list.value[editIndex.value] }
  } else {
    clearValue()
  }
}
const closeForm = () => {
  editIndex.value = -1
  formModal.value = false
}

const clearValue = () => {
  const tmp: Record<string, any> = {}
  Object.keys(value.value).forEach((k) => {
    tmp[k] = undefined
  })
  Object.assign(value.value, tmp, {})
}

const addElement = () => {
  formMod.value = 'create'
  openForm()
}

const saveElements = () => {
  props.saveHandle(list.value).then((res: boolean) => {
    if (res) {
      ElMessage.success('保存成功')
      reloadList()
      appStore.setChange()
    } else {
      ElMessage.error('保存失败')
    }
  })
}

const cb = () => {
  saveElements()
}
</script>

<style>
.cell {
  max-height: 6em;
  text-overflow: ellipsis;
}
</style>
