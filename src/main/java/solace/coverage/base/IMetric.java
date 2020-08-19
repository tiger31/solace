package solace.coverage.base;

public interface IMetric<T> {
    public float getCoverage();
    public void Reset();
    public MetricMetadata getMetadata();
    public T getSpecificationEntity();
}