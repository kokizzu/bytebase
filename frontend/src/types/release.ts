import { getProjectNameReleaseId } from "@/store/modules/v1/common";
import { EMPTY_ID, UNKNOWN_ID } from "./const";
import { Release } from "./proto/v1/release_service";
import {
  emptyProject,
  unknownProject,
  type ComposedProject,
} from "./v1/project";

export interface ComposedRelease extends Release {
  // Format: projects/{project}
  project: string;
  projectEntity: ComposedProject;
}

export const UNKNOWN_RELEASE_NAME = `${unknownProject().name}/releases/${UNKNOWN_ID}`;

export const emptyRelease = (): ComposedRelease => {
  const projectEntity = emptyProject();
  const release = Release.fromJSON({
    name: `${projectEntity.name}/releases/${EMPTY_ID}`,
    uid: String(EMPTY_ID),
    project: projectEntity.name,
  });
  return {
    ...release,
    project: projectEntity.name,
    projectEntity,
  };
};

export const unknownRelease = (): ComposedRelease => {
  const projectEntity = unknownProject();
  const release = Release.fromJSON({
    name: `${projectEntity.name}/releases/${UNKNOWN_ID}`,
    uid: String(UNKNOWN_ID),
    project: projectEntity.name,
  });
  return {
    ...release,
    project: projectEntity.name,
    projectEntity,
  };
};

export const isValidReleaseName = (name: any): name is string => {
  if (typeof name !== "string") return false;
  const [projectName, releaseName] = getProjectNameReleaseId(name);
  return Boolean(
    projectName &&
      projectName !== String(UNKNOWN_ID) &&
      releaseName &&
      releaseName !== String(UNKNOWN_ID)
  );
};
