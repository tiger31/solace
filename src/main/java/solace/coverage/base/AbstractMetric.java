package solace.coverage.base;

public abstract class AbstractMetric<T> implements IMetric<T> {
    private T specificationEntity;

    public AbstractMetric(T entity) {
        this.specificationEntity = entity;
    }

    @Override
    public T getSpecificationEntity() {
        return this.specificationEntity;
    }
}