<template>
  <div class="text-lg leading-6 font-medium text-main">
    <i18n-t keypath="repository.version-control-status">
      <template #status>
        <span class="text-success"> {{ $t("common.enabled") }} </span>
      </template>
    </i18n-t>
  </div>
  <div class="mt-2 textinfolabel">
    <i18n-t keypath="repository.version-control-description-file-path">
      <template #fullPath>
        <a class="normal-link" :href="repository.webUrl" target="_blank">{{
          repository.fullPath
        }}</a>
      </template>
      <template #fullPathTemplate>
        <span class="font-medium text-main"
          >{{ state.repositoryConfig.baseDirectory }}/{{
            state.repositoryConfig.filePathTemplate
          }}</span
        >
      </template>
    </i18n-t>
    <span>&nbsp;</span>
    <i18n-t keypath="repository.version-control-description-branch">
      <template #branch>
        <span class="font-medium text-main">
          <template v-if="state.repositoryConfig.branchFilter">
            {{ state.repositoryConfig.branchFilter }}
          </template>
          <template v-else>
            {{ $t("common.default") }}
          </template>
        </span>
      </template>
    </i18n-t>
    <template v-if="state.repositoryConfig.schemaPathTemplate">
      <span>&nbsp;</span>
      <i18n-t
        keypath="repository.version-control-description-description-schema-path"
      >
        <template #schemaPathTemplate>
          <span class="font-medium text-main">{{
            state.repositoryConfig.schemaPathTemplate
          }}</span>
        </template>
      </i18n-t>
    </template>
  </div>
  <RepositoryForm
    class="mt-4"
    :allow-edit="allowEdit"
    :vcs-type="repository.vcs.type"
    :vcs-name="repository.vcs.name"
    :repository-info="repositoryInfo"
    :repository-config="state.repositoryConfig"
    :project="project"
    :schema-migration-type="state.schemaMigrationType"
    @change-schema-migration-type="(type) => (state.schemaMigrationType = type)"
    @change-repository="$emit('change-repository')"
  />
  <div v-if="allowEdit" class="mt-4 pt-4 flex border-t justify-between">
    <BBButtonConfirm
      :style="'RESTORE'"
      :button-text="$t('repository.restore-to-ui-workflow')"
      :require-confirm="true"
      :ok-text="$t('common.restore')"
      :confirm-title="$t('repository.restore-to-ui-workflow') + '?'"
      :confirm-description="$t('repository.restore-ui-workflow-description')"
      @confirm="restoreToUIWorkflowType"
    />
    <div>
      <button
        type="button"
        class="btn-primary ml-3 inline-flex justify-center py-2 px-4"
        :disabled="!allowUpdate"
        @click.prevent="doUpdate"
      >
        {{ $t("common.update") }}
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, reactive, watch } from "vue";
import isEmpty from "lodash-es/isEmpty";
import RepositoryForm from "./RepositoryForm.vue";
import {
  Repository,
  RepositoryPatch,
  ExternalRepositoryInfo,
  RepositoryConfig,
  Project,
  ProjectPatch,
  SchemaMigrationType,
} from "../types";
import { useI18n } from "vue-i18n";
import { pushNotification, useProjectStore, useRepositoryStore } from "@/store";
import { isDev } from "@/utils";

interface LocalState {
  repositoryConfig: RepositoryConfig;
  schemaMigrationType: SchemaMigrationType;
}

export default defineComponent({
  name: "RepositoryPanel",
  components: { RepositoryForm },
  props: {
    project: {
      required: true,
      type: Object as PropType<Project>,
    },
    repository: {
      required: true,
      type: Object as PropType<Repository>,
    },
    allowEdit: {
      default: true,
      type: Boolean,
    },
  },
  emits: ["change-repository"],
  setup(props) {
    const { t } = useI18n();
    const repositoryStore = useRepositoryStore();
    const state = reactive<LocalState>({
      repositoryConfig: {
        baseDirectory: props.repository.baseDirectory,
        branchFilter: props.repository.branchFilter,
        filePathTemplate: props.repository.filePathTemplate,
        schemaPathTemplate: props.repository.schemaPathTemplate,
        sheetPathTemplate: props.repository.sheetPathTemplate,
      },
      schemaMigrationType: props.project.schemaMigrationType,
    });

    watch(
      () => props.repository,
      (cur) => {
        state.repositoryConfig = {
          baseDirectory: cur.baseDirectory,
          branchFilter: cur.branchFilter,
          filePathTemplate: cur.filePathTemplate,
          schemaPathTemplate: cur.schemaPathTemplate,
          sheetPathTemplate: cur.sheetPathTemplate,
        };
      }
    );

    const repositoryInfo = computed((): ExternalRepositoryInfo => {
      return {
        externalId: props.repository.externalId,
        name: props.repository.name,
        fullPath: props.repository.fullPath,
        webUrl: props.repository.webUrl,
      };
    });

    const allowUpdate = computed(() => {
      return (
        !isEmpty(state.repositoryConfig.branchFilter) &&
        !isEmpty(state.repositoryConfig.filePathTemplate) &&
        (props.repository.branchFilter !==
          state.repositoryConfig.branchFilter ||
          props.repository.baseDirectory !==
            state.repositoryConfig.baseDirectory ||
          props.repository.filePathTemplate !==
            state.repositoryConfig.filePathTemplate ||
          props.repository.schemaPathTemplate !==
            state.repositoryConfig.schemaPathTemplate ||
          props.repository.sheetPathTemplate !==
            state.repositoryConfig.sheetPathTemplate ||
          props.project.schemaMigrationType !== state.schemaMigrationType)
      );
    });

    const restoreToUIWorkflowType = () => {
      repositoryStore.deleteRepositoryByProjectId(props.project.id).then(() => {
        pushNotification({
          module: "bytebase",
          style: "SUCCESS",
          title: t("repository.restore-ui-workflow-success"),
        });
      });
    };

    const doUpdate = async () => {
      const repositoryPatch: RepositoryPatch = {};
      if (
        props.repository.branchFilter != state.repositoryConfig.branchFilter
      ) {
        repositoryPatch.branchFilter = state.repositoryConfig.branchFilter;
      }
      if (
        props.repository.baseDirectory != state.repositoryConfig.baseDirectory
      ) {
        repositoryPatch.baseDirectory = state.repositoryConfig.baseDirectory;
      }
      if (
        props.repository.filePathTemplate !=
        state.repositoryConfig.filePathTemplate
      ) {
        repositoryPatch.filePathTemplate =
          state.repositoryConfig.filePathTemplate;
      }
      if (
        props.repository.schemaPathTemplate !=
        state.repositoryConfig.schemaPathTemplate
      ) {
        repositoryPatch.schemaPathTemplate =
          state.repositoryConfig.schemaPathTemplate;
      }
      if (
        props.repository.sheetPathTemplate !=
        state.repositoryConfig.sheetPathTemplate
      ) {
        repositoryPatch.sheetPathTemplate =
          state.repositoryConfig.sheetPathTemplate;
      }

      // Update project schemaMigrationType field firstly.
      if (
        isDev() &&
        state.schemaMigrationType !== props.project.schemaMigrationType
      ) {
        const projectPatch: ProjectPatch = {
          schemaMigrationType: state.schemaMigrationType,
        };
        await useProjectStore().patchProject({
          projectId: props.project.id,
          projectPatch,
        });
      }

      await repositoryStore.updateRepositoryByProjectId({
        projectId: props.project.id,
        repositoryPatch,
      });
      pushNotification({
        module: "bytebase",
        style: "SUCCESS",
        title: t("repository.update-version-control-config-success"),
      });
    };

    return {
      state,
      repositoryInfo,
      allowUpdate,
      restoreToUIWorkflowType,
      doUpdate,
    };
  },
});
</script>
